#include <cassert>
#include <iostream>
 
unsigned long long gcd(unsigned long long a, unsigned long long b) {
    while (b != 0) {
        unsigned long long temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}
 
unsigned long long Phi(unsigned long long n) {
    assert(n > 0);
 
    if (n == 1)
        return 1;
 
    for (unsigned long long d : { 2, 3, 5 }) {
        if (n % d == 0) {
            unsigned long long p = 1;
            for (n /= d; n % d == 0; p *= d, n /= d);
            return (d - 1) * p * Phi(n);
        }
    }
 
    for (unsigned long long d = 7;;) {
        for (unsigned a : { 4, 2, 4, 2, 4, 6, 2, 6 }) {
            if (d * d > n)
                return n - 1;
 
            if (n % d == 0) {
                unsigned long long p = 1;
                for (n /= d; n % d == 0; p *= d, n /= d);
                return (d - 1) * p * Phi(n);
            }
 
            d += a;
        }
    }
 
    assert(false); // unreachable
}
 
int main() {
    unsigned long long k = 0;
    std::cin >> k;
 
    assert(k >= 1);
    --k;
 
    unsigned long long n = 0, d = 1;
    if (k > 0) {
        --k;
 
        unsigned long long sum_Phi = 0, next_sum_Phi;
        for (d = 1;; sum_Phi = next_sum_Phi, ++d) {
            next_sum_Phi = sum_Phi + Phi(d);
            if (next_sum_Phi > k)
                break;
        }
 
        k -= sum_Phi;
 
        for (n = 1;; ++n) {
            if (n == d) {
                std::cout << "1/1" << std::endl;
                return 0;
            }
            if (gcd(n, d) == 1 && k-- == 0)
                break;
        }
    }
 
    assert(n < d && gcd(n, d) == 1);
    std::cout << n << "/" << d << std::endl;
 
    return 0;
}