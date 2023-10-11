# JSON Sum Calculator

This program calculates the sum of numbers from a JSON file containing an array of objects. Each object has two properties, 'a' and 'b', representing the numbers to be summed.

## Input

The input JSON file should be in the following format:

```json
[
    {"a": 1, "b": 3},
    {"a": 5, "b": -9},
    {"a": -2, "b": 4},
]
```
The number of objects in the array is 1,000,000, and the values of 'a' and 'b' range from -10 to 10.

## Usage
```bash
./json_sum_calculator -file data.json -gc 5
```
- `-goroutines` (optional): Number of goroutines to use for parallel processing. Default is 1.

## Output
The program will output the total sum of all numbers to the console.

## Building and Running
1. Clone the repository:
```bash
   git clone https://github.com/atabayev/mechta.git
```
2. Build the program:
```bash
  cd json-sum-calculator
  go build
```
3. Run the program:
```bash
   ./json_sum_calculator -file data.json -gc 4
```

## Notes
    - The program reads the JSON file, processes it using the specified number of goroutines, and prints the total sum to the console.
    - For optimal performance, adjust the number of goroutines based on your system's capabilities.
    - Make sure to provide the correct path to your input JSON file.

## Repository
The code for this program can be found in the [GitHub repository](https://github.com/atabayev/mechta). Feel free to access it for more details.

## License
This program is open-source and available under the [MIT License](https://en.wikipedia.org/wiki/MIT_License). Feel free to use and modify it as needed.
