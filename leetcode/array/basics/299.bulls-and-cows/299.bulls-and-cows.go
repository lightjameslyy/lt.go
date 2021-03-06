package main

/*
 * @lc app=leetcode id=299 lang=golang
 *
 * [299] Bulls and Cows
 *
 * https://leetcode.com/problems/bulls-and-cows/description/
 *
 * algorithms
 * Medium (38.22%)
 * Total Accepted:    86.5K
 * Total Submissions: 226.4K
 * Testcase Example:  '"1807"\n"7810"'
 *
 * You are playing the following Bulls and Cows game with your friend: You
 * write down a number and ask your friend to guess what the number is. Each
 * time your friend makes a guess, you provide a hint that indicates how many
 * digits in said guess match your secret number exactly in both digit and
 * position (called "bulls") and how many digits match the secret number but
 * locate in the wrong position (called "cows"). Your friend will use
 * successive guesses and hints to eventually derive the secret number.
 *
 * Write a function to return a hint according to the secret number and
 * friend's guess, use A to indicate the bulls and B to indicate the cows.
 *
 * Please note that both secret number and friend's guess may contain duplicate
 * digits.
 *
 * Example 1:
 *
 *
 * Input: secret = "1807", guess = "7810"
 *
 * Output: "1A3B"
 *
 * Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.
 *
 * Example 2:
 *
 *
 * Input: secret = "1123", guess = "0111"
 *
 * Output: "1A1B"
 *
 * Explanation: The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a
 * cow.
 *
 * Note: You may assume that the secret number and your friend's guess only
 * contain digits, and their lengths are always equal.
 */

/*
 * using hash
 *

func getHint1(secret string, guess string) string {
    bulls, cows := 0, 0
    secretMap := make(map[byte]int)
    guessMap := make(map[byte]int)
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            if v, ok := secretMap[secret[i]]; ok {
                secretMap[secret[i]] = v+1
            } else {
                secretMap[secret[i]] = 1
            }
            if v, ok := guessMap[guess[i]]; ok {
                guessMap[guess[i]] = v+1
            } else {
                guessMap[guess[i]] = 1
            }
        }
    }
    for k, v1 := range secretMap {
        if v2, ok := guessMap[k]; ok {
            cows = cows + func(a, b int) int {
                if a <= b {
                    return a
                }
                return b
            }(v1, v2)
        }
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)
}

*/

/* using less space in hash */
func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	secretMap := make([]int, 10)
	guessMap := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretMap[int(secret[i])-'0']++
			guessMap[int(guess[i])-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		if secretMap[i] <= guessMap[i] {
			cows = cows + secretMap[i]
		} else {
			cows = cows + guessMap[i]
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
