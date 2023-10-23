# Employees - Nearest City Assignment

"Employees," is a simple Go program that its story is about a company assign its employees to the nearest city for a business mission based on their seniority year in the company. The mission has a capacity constraint for three destination cities: a, b, and c. The distances from the origin cities (Munich, Hamburg, and Berlin) to these destination cities are used to make the assignment.

## City Capacities

- City a: 2 employees
- City b: 1 employee
- City c: 2 employees

## Distance Table

This table represents the distances (in kilometers) from Munich, Hamburg, and Berlin to the destination cities a, b, and c:

| (Km)    | a   | b   | c   |
|---------|-----|-----|-----|
| Munich  | 200 | 250 | 300 |
| Hamburg | 400 | 150 | 200 |
| Berlin  | 100 | 200 | 300 |

## How It Works

The program stores information about employees, including their name, city of origin, and date of joining the company. It then calculates the nearest city for each employee based on their joining date and the capacity constraints for the destination cities.

## Test Case

Here is the test case and the expected output for the provided employee information:

- Eric:
    - Name: Eric
    - City: Berlin
    - Start Time: 2012

- John:
    - Name: John
    - City: Munich
    - Start Time: 2010

- Nami:
    - Name: Nami
    - City: Hamburg
    - Start Time: 2009

- Anthony:
    - Name: Anthony
    - City: Hamburg
    - Start Time: 2021

- Jeniffer:
    - Name: Jeniffer
    - City: Munich
    - Start Time: 2009
  
Output:

1 ) Name:  Eric,  City:  Berlin,  StartTime:  2012,  DepartCity:  c

2 ) Name:  John,  City:  Munich,  StartTime:  2010,  DepartCity:  a

3 ) Name:  Nami,  City:  Hamburg,  StartTime:  2009,  DepartCity:  b

4 ) Name:  Anthony,  City:  Hamburg,  StartTime:  2021,  DepartCity:  c

5 ) Name:  Jeniffer,  City:  Munich,  StartTime:  2009,  DepartCity:  a


## How to Use

1. Clone this repository:
   ```bash
   git clone https://github.com/A-Eric-J/Employees.git
   ```

2. Build and Run the Go program.


3. You can add more employees and modify their information in the `employees.go` file.

## Contributions

Contributions to this project are welcome. If you have any suggestions, improvements, or feature requests, please open an issue or create a pull request.
