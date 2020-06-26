package io.github.raqbit.aoc2k16;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Created by Raqbit on 1-12-16.
 */
public class Day1 {

    static Coord coord = new Coord(0, 0);
    static List<Coord> coordList = new ArrayList<Coord>();

    static Coord doubleCoord = null;

    // facing, 0 = north, 1= east, 2 = south & 3 = west
    static int direction = 0;

    public static void exec() {

        String input = "";

        //Loading file
        Scanner scanner = null;
        try {
            scanner = new Scanner(new File("day1.input"));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        input = input + scanner.nextLine();

        scanner.close();

        // route
        final String[] parsedInput = input.split(", ");

        // For every instruction
        for (String s : parsedInput) {
            if (s.startsWith("R")) {
                changeCoord(true, Integer.parseInt(s.substring(1)));
                direction++;
            } else {
                changeCoord(false, Integer.parseInt(s.substring(1)));
                direction--;
            }


            fixDirections();

            System.out.println("Instruction: " + s);
            System.out.println("--------------");
            System.out.println("X: " + coord.x);
            System.out.println("Y: " + coord.y);
            System.out.println("Direction: " + direction);
            System.out.println("--------------");
        }

        System.out.println("");
        System.out.println("Shortest route is: " + getShortestRoute());
    }

    static int getShortestRoute() {
        return Math.abs(0 - coord.x) + Math.abs(0 - coord.y);
    }

    static void changeCoord(boolean right, int amount) {
        // North
        if (direction == 0) {
            if (right) {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x + i, coord.y);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            } else {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x - i, coord.y);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            }
        }
        // East
        else if (direction == 1) {
            if (right) {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x, coord.y - i);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            } else {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x, coord.y + i);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;

                }
            }
        }
        // South
        else if (direction == 2) {
            if (right) {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x - i, coord.y);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;

                }
            } else {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x + i, coord.y);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            }
        }
        // West
        else if (direction == 3) {
            if (right) {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x, coord.y + i);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            } else {
                for (int i = 1; i <= amount; i++) {
                    Coord tempCoord = new Coord(coord.x, coord.y - i);
                    checkIfContainsCoord(tempCoord);
                    coord = tempCoord;
                }
            }
        }
    }

    static void checkIfContainsCoord(Coord coord) {
        if (doubleCoord == null) {
            for (Coord listcoord : coordList) {
                if (listcoord == coord)
                    doubleCoord = coord;
            }
            if (doubleCoord == null)
                coordList.add(coord);
        } else {
            coordList.add(coord);
        }
    }

    static void fixDirections() {
        if (direction == 4) {
            direction = 0;
        } else if (direction == -1) {
            direction = 3;
        }
    }
}
