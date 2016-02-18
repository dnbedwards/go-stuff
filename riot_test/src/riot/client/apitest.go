package main

import (
	"os"
	"fmt"
	"github.com/jmcvetta/restclient"
)

type SummonerResponse struct { 
	Id string
}

const apiKey string = "46b1e549-64ec-42f9-9f91-6334d86db502"

func main() {
	summonerName := os.Args[1]
	fmt.Println("Searching summoner", summonerName)

	summonerResp := SummonerResponse {}

	response := restclient.RequestResponse {
		Url: fmt.Sprintf("https://oce.api.pvp.net/api/lol/oce/v1.4/summoner/by-name/%s?api_key=%s", summonerName, apiKey),
		Method: restclient.GET,
		Result: &summonerResp,
	}
	fmt.Println("Summoner ID is", summonerResp.Id)
}

