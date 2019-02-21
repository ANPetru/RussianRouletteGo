package main

import( "bufio"
		"strings"
		"fmt"
		"os"
		"math/rand"
		"time")

//Player asd
type Player struct{
	name string
	dead bool
}

var chamber = [6]int{1,2,3,4,5,6} 
var players []Player
var numPlayers int 
var currentPlayerIndex = -1
var bulletIndex int 
var gameOver = false

func main(){
	bulletIndex = generateBulletIndex()
	getNumberOfPlayers()
	initPlayers()
	fmt.Println("------------------------------------------------------")
	fmt.Println("Starting game")
	fmt.Println("------------------------------------------------------")

	for !gameOver {
		displayChamber()
		getCurrentPlayer()
		playTurn()
	}
}

func playTurn(){
	buf := bufio.NewReader(os.Stdin)
	hit := true
	for hit{
		hit = false
		fmt.Printf("%s's turn, spin the chamber ('spin') or quit ('quit')",players[currentPlayerIndex].name)
		str, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			str = strings.Replace(str, "\n", "", -1)
			str = getNameFromString(str)

			if str == "spin"{
				calculateAndDisplayBullet()	
			} else if str == "quit"{
				fmt.Println(players[currentPlayerIndex].name + " has left the game")
				players[currentPlayerIndex].dead = true
				checkWinner()
			} else {
				hit = true
			}
		}
	}

}

func calculateAndDisplayBullet(){
	tmpBulletIndex := generateBulletIndex() 
	for i:=0;i<3;i++{
		fmt.Print(". ")
		time.Sleep(time.Microsecond*500)
	}
	if tmpBulletIndex == bulletIndex{
		fmt.Println("Bang!")
		killCurrentPlayerAndRestartChamber()
	} else {
		chamber[tmpBulletIndex] = -1
		fmt.Println("safe")
	}

}

func killCurrentPlayerAndRestartChamber(){
	fmt.Println(players[currentPlayerIndex].name + " died")
	players[currentPlayerIndex].dead = true
	checkWinner()
	if !gameOver{
		fmt.Println("Restarting chamber")
		for i := range chamber{
			chamber[i] = i +1
		}
		bulletIndex = generateBulletIndex()
	}

}

func checkWinner(){
	//TODO: make players alive global
	playersAlive:= numPlayers
	for i:=range players{
		if players[i].dead{
			playersAlive--
		}
	}
	if playersAlive ==1{
		endGame()
	}
}

func endGame(){
	gameOver = true
	getCurrentPlayer()
	fmt.Println(players[currentPlayerIndex].name + " has won the game.\nShutting down")
}

func generateBulletIndex()int{
	hit := true
	randN := 0
	for hit{
		hit = false
		rand.Seed(time.Now().UnixNano())
		randN = rand.Intn(5)
		if chamber[randN] == -1{
			hit = true
		}
	}
	return randN
}

func displayChamber(){
	fmt.Print("Chamber:[ ")
	for i := 0; i<len(chamber);i++{
		if chamber[i] == -1{
			fmt.Print(" - ")
		} else {
			fmt.Printf(" %d ", chamber[i])
		}
	}
	fmt.Println("]")
}

func getNumberOfPlayers(){
	numPlayers=0
	for numPlayers < 2 || numPlayers >6{
		fmt.Println("Enter number of players (2-6)")
		fmt.Scanf("%d\n",&numPlayers)
		if numPlayers < 2 ||numPlayers >6{
			fmt.Println("The number of players can only be from 2 to 6")
		}
	}
}

func getCurrentPlayer(){
	if currentPlayerIndex == -1 || currentPlayerIndex == (numPlayers -1) {
		currentPlayerIndex = 0
	} else {
		hit := true
		for hit{
			hit = false
			currentPlayerIndex ++;
			if currentPlayerIndex >numPlayers -1{
				currentPlayerIndex = 0
			}
			if players[currentPlayerIndex].dead{
				hit = true
			}
		}
	}
}

func initPlayers(){
	buf := bufio.NewReader(os.Stdin)
	players = make([]Player, numPlayers)
	for i:= range players{

		fmt.Printf("Enter name for player %d\n", (i +1))
		nName, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			nName = getNameFromString(nName)
			players[i].name = nName
			players[i].dead = false
		}
	}
}

func getNameFromString(str string) string{
	name:=""
	for i := range str{
		if (str[i] >64 && str[i]<91) || (str[i]>96 && str[i]<123){
			name += string(str[i])
		}
	}
	return name
}