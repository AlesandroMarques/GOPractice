package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"

)

func main() {


	if len(os.Args) < 4{
		log.Fatalf("Usage: %s <user> <host:port> <pw> ",  os.Args[0])




	}
	ubu107:= serverSSH{os.Args[2],os.Args[1],os.Args[3] }


	client, session, err := connectToHostStruct(ubu107)
	if err != nil {
		panic(err)
	}

	//commandSlice := []string{"cd /opt/IBM/MDM","mkdir GOGO"}
	// 1 lond command things after && are not run if others fail
	// if ; instead of && still run
        longCommand := " cd /opt/IBM/MDM && mkdir GOGO && cd GOGO && echo hello > Hello.txt"

/*
	out, err := session.CombinedOutput(os.Args[4])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	out1, err1 := session.CombinedOutput("echo hello")
	if err != nil {
		panic(err1)
	}
	fmt.Println(string(out1))*/
	//runcommands(&session, os.Args[4])
	 runcommandsCombined(session, longCommand)
	// commandPipe(session, commandSlice)
fmt.Println("closing client ")
	client.Close()
}



type serverSSH struct {
	host string
	user string
	pw string
}


// adds commands to pipe but weird
func commandPipe(sesh *ssh.Session, slice []string){
	stdinBuf, _ :=sesh.StdinPipe()

	err1 := sesh.Shell()
	for _,cmd := range slice{
	   _, err:=stdinBuf.Write([]byte(cmd))
		if err !=nil{
			panic(err)
		}
		fmt.Println(cmd , "entered into bef")
	}
	if err1 != nil {
		panic(err1)
	}

fmt.Println("all commands run successfully")
}

// runs commands using RUN
func runcommands ( sesh *ssh.Session, cmd string){
	err := sesh.Run(cmd)
	if err != nil {
		panic(err)

	}
	fmt.Println("command" ,cmd, "successfull")




}
// runs with combined output
func runcommandsCombined ( sesh *ssh.Session, cmd string){
	out,err := sesh.CombinedOutput(cmd)
	if err != nil {
		panic(err)

	}
	fmt.Println("command" ,cmd, "successfull")
	fmt.Println(out)



}



func connectToHostStruct(s serverSSH) (*ssh.Client, *ssh.Session, error) {
	fmt.Println(s.user)
	fmt.Println(s.host)
	fmt.Println(s.pw)

	sshConfig := &ssh.ClientConfig{
		User: s.user,
		Auth: []ssh.AuthMethod{ssh.Password(s.pw)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

	}

	client, err := ssh.Dial("tcp", s.host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}



func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	var pass string
	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}