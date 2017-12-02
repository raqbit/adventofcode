package it.raqb.adventofcode

import com.andreapivetta.kolor.green
import com.andreapivetta.kolor.yellow
import it.raqb.adventofcode.day1.Day1
import it.raqb.adventofcode.day2.Day2

val year = "2017"

val days = arrayOf(Day1, Day2)

fun main(args: Array<String>) {

    println("==== AdventOfCode $year ====".green())
    for ((index, day) in days.withIndex()) {
        println("Day ${index + 1}: ${day.name}  ${getStars(day.starCount)}")
    }
    println("=== Solutions by Raqbit ===".green())
    println()

    var selectedDayNum: Int

    while (true) {
        print("Enter a day number: ")

        val inputString = readLine().orEmpty()

        try {
            selectedDayNum = inputString.toInt()
        } catch (exception: NumberFormatException) {
            println("That's not a number")
            continue
        }

        if (selectedDayNum < 1 || selectedDayNum > days.size) {
            println("That day doesn't exist (yet)")
            continue
        }

        break
    }

    val selectedDay = days[selectedDayNum - 1]

    println("Selected: Day $selectedDayNum.")
    println()

    val input = Day::class.java.getResource("/inputs/day$selectedDayNum.txt").readText(Charsets.UTF_8)

    selectedDay.exe(input)
}

fun getStars(starCount: Int): String {
    return when (starCount) {
        2 -> "*".yellow()
        1 -> "*"
        else -> ""
    }
}
