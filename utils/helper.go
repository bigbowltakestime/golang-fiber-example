package main

import "fmt"

// WelcomeMessage은 환영 메시지를 생성하는 함수입니다.
func WelcomeMessage(name string) string {
	return fmt.Sprintf("Welcome, %s!", name)
}

// FarewellMessage은 작별 메시지를 생성하는 함수입니다.
func FarewellMessage(name string) string {
	return fmt.Sprintf("Goodbye, %s. We'll miss you!", name)
}
