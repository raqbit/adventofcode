package it.raqb.adventofcode

import it.raqb.adventofcode.day1.Day1

val year = "2017"

val days = arrayOf(Day1)

fun main(args: Array<String>) {

    println("==== AdventOfCode $year ====")
    for ((index, day) in days.withIndex()) {
        println("Day ${index + 1}: ${day.name}  ${"*".repeat(day.starCount)}")
    }
    println("===========================")
    println()

    var selectedDayNum: Int

    while (true){
        print("Enter a valid day: ")

        val inputString = readLine().orEmpty()

        try {
            selectedDayNum = inputString.toInt()
        } catch (exception: NumberFormatException) {
            println("That's not a number")
            continue
        }

        if(selectedDayNum < 1 || selectedDayNum > days.size){
            println("That day doesn't exist (yet)")
            continue
        }

        break
    }

    val selectedDay = days[selectedDayNum - 1]

    println("Selected: Day $selectedDayNum.")
    println()

    selectedDay.exe()
}