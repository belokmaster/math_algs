#include <iostream>
#include <algorithm>

// Функция для вычисления наибольшего общего делителя (НОД)
int gcd(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

int main() {
    int k;
    std::cin >> k;

    int count = 0; // Счетчик для нахождения k-го элемента

    // Перебор значений d от 1 до бесконечности
    for (int d = 1; ; ++d) {
        // Перебор значений n от 0 до d
        for (int n = 0; n <= d; ++n) {
            // Проверка, является ли дробь несократимой
            if (gcd(n, d) == 1) {
                ++count;
                // Если нашли k-й элемент, выводим его и завершаем программу
                if (count == k) {
                    std::cout << n << "/" << d << std::endl;
                    return 0;
                }
            }
        }
    }

    return 0;
}
