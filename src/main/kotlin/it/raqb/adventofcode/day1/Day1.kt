package it.raqb.adventofcode.day1

import it.raqb.adventofcode.Day

object Day1 : Day {
    override val name = "Inverse Captcha"
    override val starCount = 2

    override fun exe(input: String) {
        part1(input)
        part2(input)
    }

    private fun part1(input: String) {
        var sum = 0

        // Loop over every character
        for ((index, char) in input.withIndex()) {
            // toInt() directly gets me the character code
            val charVal = char.toString().toInt()
            // index is equal to the next one unless it should wrap
            val checkIndex = if (index == input.length - 1) 0 else index + 1
            if (char == input[checkIndex]) {
                // Adding to the sum
                sum += charVal
            }
        }

        println("Answer to part 1: $sum.")
    }

    private fun part2(input: String) {
        var sum = 0

        var halfwayAroundSteps = input.length / 2

        // Loop over every character
        for ((index, char) in input.withIndex()) {
            // toInt() directly gets me the character code
            val charVal = char.toString().toInt()

            // Calculate new index with halfwayaroundsteps
            var checkIndex = index + halfwayAroundSteps

            // Calculate overshoot
            val overshoot = checkIndex + 1 - input.length

            // If there's an actual overshoot, that should be the new index to check
            if (overshoot > 0) {
                checkIndex = overshoot - 1
            }

            if (char == input[checkIndex]) {
                // Adding to the sum
                sum += charVal
            }
        }

        println("Answer to part 2: $sum.")
    }

}
