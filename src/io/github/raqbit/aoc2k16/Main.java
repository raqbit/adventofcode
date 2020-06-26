package io.github.raqbit.aoc2k16;

import java.util.Scanner;

/**
 * Created by Raqbit on 1-12-16.
 */
public class Main {

    public static void main(String[] args) {

        System.out.println("");
        System.out.println("Advent of Code 2016 solutions by Raqbit");
        System.out.println("");
        System.out.println("-------------------------------------------");
        System.out.println("");
        System.out.println("Please enter the number of the day to view.");
        System.out.println("");
        System.out.println("-------------------------------------------");
        System.out.println("");

        int choice = getChoice();

        switch (choice) {
            case 1: {
                Day1.exec();
                break;
            }
            case 2: {
                Day2.exec();
                break;
            }
            default: {
                System.out.println("");
                System.out.println("Well... this is akward....");
                break;
            }
        }
    }

    private static int getChoice() {
        Scanner input = new Scanner(System.in);
        int choice = 0;

        while (choice == 0) {
            try {
                choice = Integer.parseInt(input.nextLine());
            } catch (NumberFormatException exception) {
                System.out.println("Please enter a number.");
            }
        }
        input.close();
        return choice;
    }
}
