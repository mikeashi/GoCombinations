# GoCombinations

This repository contains a benchmark for different implementations of generating combinations of a given alphabet to a given length. The main goal of this project is to explore various approaches in Go.

## Installation

To use this benchmark, you need to have Go installed on your machine. You can download and install Go from the official website: https://golang.org/

## Usage

1. Clone the repository:

    ```shell
    git clone https://github.com/mikeashi/GoCombinations.git
    ```

2. Change to the project directory:

    ```shell
    cd GoCombinations
    ```

3. Run the benchmark:

    ```shell
    go run .
    ```


 ## Results from My Machine
 
    ```shell
        Running benchmarks:
        Min: 1 , Max: 4 
        Alphabet:  abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890
        +----------------+---------------+----------------+--------------+
        | Implementation | Duration      | Rate           | Combinations |
        +----------------+---------------+----------------+--------------+
        | Closure        | 2.848125231s  | 5273142.429459 | 15018570     |
        | Encoder        | 4.197992291s  | 3577560.166606 | 15018570     |
        | GPTGenerator   | 2.856906116s  | 5256935.086487 | 15018570     |
        | Iterium        | 10.544482895s | 1424305.975888 | 15018570     |
        +----------------+---------------+----------------+--------------+
    ```
## Contributing

Contributions are welcome! If you have any ideas or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.