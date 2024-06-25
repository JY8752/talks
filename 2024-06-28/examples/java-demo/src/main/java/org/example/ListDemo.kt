package org.example

enum class Status{
    Success,
    Failed
    ;
}

class FizzBuzz {
    fun run() {
        for(i in 1..15) {
            if(i % 15 == 0) {
                println("FizzBuzz!!")
            } else if(i % 5 == 0) {
                println("Buzz!!")
            } else if(i % 3 == 0) {
                println("Fizz!!")
            }
        }
    }
}

fun main() {

}