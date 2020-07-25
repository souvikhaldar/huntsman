package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"huntsman/config"
	mongodb "huntsman/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
	Hits        int32   `json:"hits"`
}

var tcpDump string
var ipAdd string
var readStream bool
var conf string
var persist bool

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringVarP(&tcpDump, "tcp-dump", "t", "", "source file of tcpdump")
	ipCmd.Flags().StringVar(&ipAdd, "ip", "", "IP address of the target")
	ipCmd.Flags().BoolVarP(&readStream, "read-stream", "s", false, "Do you want to read from tcpdump output contineously?")

	ipCmd.PersistentFlags().StringVarP(&conf, "config", "c", "", "The path to the configuration JSON file")
	ipCmd.Flags().BoolVarP(&persist, "persist", "p", false, "Do you want to store the response to mongo? If yes, please provide value to --config flag")
}

var ipCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "Fetch the location information of the IP",
	Run: func(cmd *cobra.Command, args []string) {
		var con config.Config
		if persist {
			con = config.SetEnv(conf)
			if err := mongodb.InitializeMongoDB(con); err != nil {
				log.Fatal(err)
			}
		}

		ipCache := make(map[string]*Response)
		if readStream {
			cmd := exec.Command("sh", "-c", "sudo tcpdump -s 0 -A 'tcp[((tcp[12:1] & 0xf0) >> 2):4] = 0x47455420'")

			stdOut, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Println(err)
				return
			}
			scanner := bufio.NewScanner(stdOut)
			go func() {
				for scanner.Scan() {

					if strings.Contains(scanner.Text(), "IP") {
						ip, err := ParseIPFromTcpDump(scanner.Text())
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println("Request came from: ", ip)
						if cacheRes, ok := ipCache[ip]; ok {
							fmt.Println("Cache hit")
							cacheRes.Hits += 1
							fmt.Printf(
								"Details of the IP:\n %+v \n",
								cacheRes)
							continue
						}
						body, err := getIPInfo(ip)
						if err != nil {
							fmt.Println(err)
							return
						}
						// insert to cache
						ipCache[ip] = &body
						fmt.Printf(
							"Details of the IP:\n %+v \n",
							body)
					}
					fmt.Println("------------------------")
				}
			}()
			if err := cmd.Start(); err != nil {
				log.Fatal(err)
			}
			if err := cmd.Wait(); err != nil {
				log.Fatal(err)
			}

			return
		}

		if tcpDump != "" {
			f, err := os.Open(tcpDump)
			if err != nil {
				fmt.Println("Unable to open log file", err)
				return
			}

			scanner := bufio.NewScanner(f)
			for {
				for scanner.Scan() {
					if strings.Contains(scanner.Text(), "IP") {
						ip, err := ParseIPFromTcpDump(scanner.Text())
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println("Request came from: ", ip)

						if cacheRes, ok := ipCache[ip]; ok {
							cacheRes.Hits += 1
							fmt.Printf(
								"Details of the IP:\n %+v \n",
								cacheRes)
							continue
						}
						response, err := getIPInfo(ip)
						if err != nil {
							fmt.Println(err)
							return
						}
						ipCache[ip] = &response

						fmt.Printf(
							"Details of the IP:\n %+v \n",
							response)

						if !persist {
							return
						}
						_, err = mongodb.MongoIPCollection.UpdateOne(
							context.TODO(),
							bson.D{
								{"query", response.Query},
							},
							bson.D{
								{
									"$inc", bson.D{
										{"hits", 1},
									},
								},
							},
						)
						if err != nil {
							fmt.Println("Could not update: ", err)
							return
						}
						fmt.Println("Updated to mongo")

					}
				}
				if err := scanner.Err(); err != nil {
					fmt.Println(err)
					return
				}
			}
			return
		}

		// plain IP passing
		response, err := getIPInfo(ipAdd)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Details of the IP:\n %+v \n", response)
		//if insertRes, err := mongodb.MongoIPCollection.InsertOne(
		//	context.TODO(),
		//	response,
		//); err != nil {
		//	fmt.Println("Error in inserting to mongo: ", err)
		//} else {
		//	fmt.Println("Insert ID: ", insertRes.InsertedID)
		//}
		if !persist {
			return
		}
		_, err = mongodb.MongoIPCollection.UpdateOne(
			context.TODO(),
			bson.D{
				{"query", response.Query},
			},
			bson.D{
				{
					"$inc", bson.D{
						{"hits", 1},
					},
				},
			},
		)
		if err != nil {
			fmt.Println("Could not update: ", err)
			return
		}
		fmt.Println("Updated to mongo")

	},
}

func getIPInfo(ip string) (Response, error) {
	var response Response
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error in unmarshalling response: ", err)
		return response, err
	}
	return response, err
}

func ParseIPFromTcpDump(tcpDump string) (string, error) {
	split := strings.Split(tcpDump, "\n")
	if len(split) == 0 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "split")
	}
	newSplit := strings.SplitAfter(split[0], ">")[0]
	sliceSplit := strings.Fields(newSplit)
	if len(sliceSplit) < 3 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "sliceSplit")
	}
	lastDot := strings.LastIndex(sliceSplit[2], ".")
	if lastDot == -1 {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "lastDot")
	}
	if len(sliceSplit[2]) < lastDot {
		return "", fmt.Errorf("Error in parsing tcpdump output: %s", "lastDot length")
	}
	return sliceSplit[2][:lastDot], nil
}
