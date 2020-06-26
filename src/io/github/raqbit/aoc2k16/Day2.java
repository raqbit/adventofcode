package io.github.raqbit.aoc2k16;


import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Created by Raqbit on 1-12-16.
 */
public class Day2 {

    //String list of
    static final List<String> instructionFile = new ArrayList<>();

    // Keypad used to open the bathroom
    static final int[][] keypad = {
            {7, 8, 9},
            {4, 5, 6},
            {1, 2, 3}};

    // Last number that was put in the password
    static int lastNumber;

    // Final password
    static String password = "";

    // Current coordinate;
    static Coord coord = new Coord(1, 1);

    public static void exec() {

        //Loading file
        Scanner s = null;
        try {
            s = new Scanner(new File("day2.input"));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        while (s.hasNextLine()) {
            instructionFile.add(s.nextLine());
        }
        s.close();

        for (String line : instructionFile) {
            for (char direction : line.toCharArray()) {
                switch (direction) {
                    case 'U': {
                        up();
                        break;
                    }
                    case 'D': {
                        down();
                        break;
                    }
                    case 'L': {
                        left();
                        break;
                    }
                    case 'R': {
                        right();
                        break;
                    }
                    default: {
                        System.out.println("Invalid character found.");
                        break;
                    }
                }
                System.out.println("Going " + direction);
                System.out.println("Coords: " + coord.x + ", " + coord.y);
                System.out.println("----------------");
            }
            addToPass(keypad[coord.y][coord.x]);
        }
        System.out.println("The final password is:");
        System.out.println(password);

    }

    private static void up() {
        if (coord.y != 2)
            coord.y += 1;

    }

    private static void down() {
        if (coord.y != 0)
            coord.y -= 1;
    }

    private static void left() {
        if (coord.x != 0)
            coord.x -= 1;
    }

    private static void right() {
        if (coord.x != 2)
            coord.x += 1;
    }

    private static void addToPass(int number) {
        if (number != lastNumber) {
            password = password + String.valueOf(number);
            lastNumber = number;
        }
    }
}
