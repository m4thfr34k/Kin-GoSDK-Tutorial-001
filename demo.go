package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kinecosystem/agora-common/kin"
	"github.com/kinecosystem/kin-go/client"
)

func main() {

	var tokaccount []kin.PublicKey     //Account on the Solana blockchain
	var desttokaccount []kin.PublicKey //Account on the Solana blockchain
	var priv, dest kin.PrivateKey

	// Environments - EnvironmentProd or EnvironmentTest
	// KinVersion(4) = the current blockchain / Solana
	c, err := client.New(client.EnvironmentTest, client.WithKinVersion(4))
	if err != nil {
		log.Fatal(err)
	}

	// Setting up the priv account - PrivateKey -> CreateAccount -> ResolveTokenAccount
	priv, err = kin.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	err = c.CreateAccount(context.Background(), priv)
	if err != nil {
		log.Fatal(err)
	}

	// prints the address in a format that can be used on the Solana explorer
	fmt.Println("Priv Address ->", priv.Public().Base58())

	//Kin accounts on Solana consist of the Solana account itself and the token account
	tokaccount, err = c.ResolveTokenAccounts(context.Background(), priv.Public())
	if err != nil {
		log.Fatal(err)
	}

	// Setting up the dest account
	dest, err = kin.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	err = c.CreateAccount(context.Background(), dest)
	if err != nil {
		log.Fatal(err)
	}

	// prints the address in a format that can be used on the Solana explorer
	fmt.Println("Dest Address ->", dest.Public().Base58())

	//Kin accounts on Solana consist of the Solana account itself and the token account
	desttokaccount, err = c.ResolveTokenAccounts(context.Background(), dest.Public())
	if err != nil {
		log.Fatal(err)
	}

	// We are working in the TEST environment
	// Transactions are slower in TEST and we're going to wait a few seconds
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(15 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	quarks, err := c.GetBalance(context.Background(), priv.Public())
	fmt.Println("Priv Address ->", priv.Public().Base58())
	fmt.Println("Priv Quarks  ->", quarks)

	quarks, err = c.GetBalance(context.Background(), dest.Public())
	fmt.Println("Dest Address ->", dest.Public().Base58())
	fmt.Println("Dest Quarks  ->", quarks)

	// The Airdrop only works for Kin(4) and in TEST
	// Airdrop returns transaction information but we don't need it for the demo
	fmt.Println("Airdropping to Priv")
	_, err = c.RequestAirdrop(context.Background(), tokaccount[0], 110000)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(15 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	// Airdrop returns transaction information but we don't need it for the demo
	fmt.Println("Airdropping to Dest")
	_, err = c.RequestAirdrop(context.Background(), desttokaccount[0], 1111)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(15 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	// 1) Request airdrop 2) ... 3) PROFIT!
	quarks, err = c.GetBalance(context.Background(), priv.Public())
	fmt.Println("Priv Account ->", priv.Public())
	fmt.Println("Priv Address ->", priv.Public().Base58())
	fmt.Println("Priv Quarks  ->", quarks)

	quarks, err = c.GetBalance(context.Background(), dest.Public())
	fmt.Println("Dest Account ->", dest.Public())
	fmt.Println("Dest Address ->", dest.Public().Base58())
	fmt.Println("Dest Quarks  ->", quarks)

	fmt.Println("Time to submit a payment")
	// Submit payment returns transaction information but we don't need it for the demo
	_, err = c.SubmitPayment(context.Background(), client.Payment{
		Sender:      priv,
		Destination: dest.Public(),
		Type:        kin.TransactionTypeEarn, // kin.TransactionTypeP2P | kin.TransactionTypeSpend | kin.TransactionTypeEarn
		Quarks:      25,
		Memo:        "My demo payment",
	})

	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
	time.Sleep(15 * time.Second)
	fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())

	quarks, err = c.GetBalance(context.Background(), priv.Public())
	fmt.Println("Priv Address ->", priv.Public().Base58())
	fmt.Println("Priv Quarks  ->", quarks)

	quarks, err = c.GetBalance(context.Background(), dest.Public())
	fmt.Println("Dest Address ->", dest.Public().Base58())
	fmt.Println("Dest Quarks  ->", quarks)

}
