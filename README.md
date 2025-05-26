# Go Diffie-Hellman Key Exchange

This repository contains a Go implementation of the Diffie-Hellman key exchange algorithm.

## Algorithm Explanation

The Diffie-Hellman key exchange algorithm allows two parties to establish a shared secret key over an insecure communication channel. The steps involved are:

1.  **Parameter Generation:**
    *   A large prime number `p` is generated.
    *   A primitive root `g` modulo `p` is found. `p` and `g` can be made public.

2.  **Private and Public Key Generation:**
    *   **Alice** chooses a secret private key `a` (a large integer) and computes her public key `A = g^a mod p`.
    *   **Bob** chooses a secret private key `b` (a large integer) and computes his public key `B = g^b mod p`.

3.  **Public Key Exchange:**
    *   Alice sends her public key `A` to Bob.
    *   Bob sends his public key `B` to Alice.

4.  **Shared Secret Key Computation:**
    *   **Alice** computes the shared secret key: `S = B^a mod p`.
    *   **Bob** computes the shared secret key: `S = A^b mod p`.

Both Alice and Bob will arrive at the same value for `S`, which is their shared secret key. An eavesdropper who only knows `p`, `g`, `A`, and `B` cannot easily compute `S` due to the difficulty of the discrete logarithm problem.

The `main.go` file in this repository demonstrates this process.

## Project Structure

*   `main.go`: An example program that demonstrates the Diffie-Hellman key exchange by simulating the roles of Alice and Bob.
*   `utils/`: This directory contains the core cryptographic functions for the Diffie-Hellman algorithm:
    *   `GenerateLargePrime.go`: Implements the generation of a large prime number `p`.
    *   `FindPrimitiveRoot.go`: Implements the logic to find a primitive root `g` for the given prime `p`.
    *   `GeneratePrivateKey.go`: Implements the generation of private keys `a` or `b`.
    *   `GeneratePubKey.go`: Implements the calculation of public keys `A` or `B`.
    *   `GenerateSharedKey.go`: Implements the calculation of the shared secret key `S`.
*   `Makefile`: Contains instructions for building and testing the project.
*   `bin/`: This directory stores the compiled executable file (`DiffieHellman.exe`) after a successful build.
*   `LICENSE`: Contains the MIT License under which this project is released.
*   `test/`: Contains unit tests for the utility functions, for example `find_primitive_root_test.go`.

## Usage Instructions

To build and run this project:

1.  **Build the project:**
    Open your terminal, navigate to the project's root directory, and run the make command:
    ```bash
    make build
    ```
    This command will compile the `main.go` file and place the executable `DiffieHellman.exe` in the `bin/` directory.

2.  **Run the executable:**
    After a successful build, you can run the program using:
    ```bash
    ./bin/DiffieHellman.exe
    ```
    The program will then output the generated public keys for Alice (A) and Bob (B), and their computed shared keys, along with a confirmation if they match.

## License

This code is released under the MIT License. See the `LICENSE` file for details.
