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
	dead, quit bool
	round int
}

var chamber = [6]int{1,2,3,4,5,6} 
var players []Player
var numPlayers int 
var currentPlayerIndex int
var bulletIndex int 
var gameOver bool
var rounds int

func main(){
	playGame()
}

func playGame(){
	gameOver = false
	rounds = 1
	currentPlayerIndex = -1
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
	showEndGameOptions()
}

func showEndGameOptions(){
	fmt.Println("------------------------------------------------------")
	fmt.Println("'1'-Play Again")
	fmt.Println("'2'-Show last game stats")
	fmt.Println("'3'-Quit")
	buf := bufio.NewReader(os.Stdin)
	str, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		str = string(str[0])
		switch str{
		case "1":
			playGame()
		case "2":
			printStats()
		case "3":
			return
		default:
			showEndGameOptions()
		}
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
			str = getNameFromString(str)

			if str == "spin"{
				calculateAndDisplayBullet()	
			} else if str == "quit"{
				fmt.Println(players[currentPlayerIndex].name + " has left the game")
				players[currentPlayerIndex].dead = true
				players[currentPlayerIndex].quit = true
				players[currentPlayerIndex].round = rounds
				
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
		time.Sleep(time.Millisecond*500)
	}
	if tmpBulletIndex == bulletIndex{
		fmt.Println("Bang!")
		time.Sleep(time.Millisecond*500)

		killCurrentPlayerAndRestartChamber()
	} else {
		chamber[tmpBulletIndex] = -1
		fmt.Println("safe")
		time.Sleep(time.Millisecond*500)

	}

}

func killCurrentPlayerAndRestartChamber(){
	fmt.Println(players[currentPlayerIndex].name + " died")
	players[currentPlayerIndex].dead = true
	players[currentPlayerIndex].round = rounds
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
	fmt.Println(players[currentPlayerIndex].name + " has won the game.")

	
}

func printStats(){
	fmt.Println("------------------------------------------------------")
	fmt.Println("Name\tStat\tRound")
	for i := range players{
		fmt.Print(players[i].name)
		if players[i].quit {
			fmt.Printf("\tquit\t%d\n", players[i].round)
		} else if players[i].dead{
			fmt.Printf("\tdied\t%d\n", players[i].round)
		} else {
			fmt.Println("\twon")
		}
	}
	fmt.Println("------------------------------------------------------")
	showEndGameOptions()
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
	if currentPlayerIndex == -1 {
		currentPlayerIndex = 0
	} else{
		if currentPlayerIndex == (numPlayers -1){
			rounds++
			currentPlayerIndex = 0
		} else{
			hit := true
			for hit{
				hit = false
				if currentPlayerIndex == (numPlayers -1){
					rounds++
					currentPlayerIndex = 0
				} else {
					currentPlayerIndex ++;
				}
				if players[currentPlayerIndex].dead{
					hit = true
				}
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
			players[i].quit = false
		}
	}
}

func getNameFromString(str string) string{
	str = strings.Replace(str, "\n", "", -1)
	name:=""
	for i := range str{
		if (str[i] >64 && str[i]<91) || (str[i]>96 && str[i]<123){
			name += string(str[i])
		}
	}
	return name
}