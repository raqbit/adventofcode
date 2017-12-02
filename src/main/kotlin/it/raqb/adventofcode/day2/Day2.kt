package it.raqb.adventofcode.day2

import it.raqb.adventofcode.Day

object Day2 : Day {

    override val name = "Corruption Checksum"
    override val starCount = 2

    override fun exe(input: String) {
        val processedInput = inputToArray(input)
        part1(processedInput)
        part2(processedInput)
    }

    fun inputToArray(input: String): ArrayList<Array<Int>> {
        val values = arrayListOf<Array<Int>>()
        // Seperate lines by newline
        val lineArray = input.split(System.lineSeparator())
        // For each line
        lineArray.forEach {
            // Splitting by any space or tab, then creating ints for each item
            val intArray = it.split(Regex("\\s")).map { it.toInt() }.toTypedArray()
            values.add(intArray)
        }

        return values
    }

    fun part1(input: ArrayList<Array<Int>>) {
        var sum = 0

        // Looping over every line
        input.forEach {
            // Max and min shouldn't ever be null
            sum += it.max()!! - it.min()!!
        }

        println("The answer to part 1 is $sum.")
    }

    fun part2(input: ArrayList<Array<Int>>) {
        var sum = 0

        // Looping over every row
        for (row in input) {
            // Looping over every number
            for ((currNumIndex, currNum) in row.withIndex()) {
                row
                        .filterIndexed { checkNumIndex, checkNum -> currNum % checkNum == 0 && currNumIndex != checkNumIndex }
                        .forEach { sum += currNum / it }
            }
        }

        println("The answer to part 2 is $sum.")
    }
}