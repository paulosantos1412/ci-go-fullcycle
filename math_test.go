package main

import "testing"


func TestMain(t *testing.T){
		main()
}

func TestSum(t *testing.T){
	
	result := sum(2,3)

	if result !=5 {
		t.Error("The result must 5!")
	}
}

func TestSub(t *testing.T){
	
	result := sub(8,3)

	if result !=5 {
		t.Error("The result must 5!")
	}
}

func TestTimes(t *testing.T){
	
	result := times(2,3)

	if result !=6 {
		t.Error("The result must 5!")
	}
}


func TestDivision(t *testing.T){
	
	result := division(12,3)

	if result !=4 {
		t.Error("The result must 5!")
	}
}