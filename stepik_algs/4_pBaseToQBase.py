# Функция для преобразования числа из системы счисления с основанием p в систему с основанием q
def pBaseToQBase(num, p, q):
    num = str(num)
    # Преобразуем число из системы счисления p в массив цифр
    digits = [char_to_int(ch) for ch in num[::-1]]

    # Преобразуем в систему счисления q
    result = []
    while digits:
        # Получаем остаток от деления на q
        remainder = 0
        for i in range(len(digits) - 1, -1, -1):
            # Умножаем текущую цифру на основание p
            digits[i] += remainder * p
            # Получаем новую цифру и остаток от деления на q
            remainder = digits[i] % q
            digits[i] //= q

        # Добавляем остаток в результат
        if remainder < 10:
            result.append(chr(remainder + ord('0')))
        else:
            result.append(chr(remainder - 10 + ord('A')))

        # Удаляем ведущие нули
        while digits and digits[-1] == 0:
            digits.pop()

    # Переворачиваем результат, так как он был построен в обратном порядке
    if not result:
        return "0"

    result.reverse()
    return ''.join(result)

# Преобразование символа в целое число
def char_to_int(ch):
    if '0' <= ch <= '9':
        return ord(ch) - ord('0')
    elif 'A' <= ch <= 'Z':
        return ord(ch) - ord('A') + 10
    elif 'a' <= ch <= 'z':
        return ord(ch) - ord('a') + 10
    else:
        raise ValueError(f"Invalid character: {ch}")


num, a, b = map(int, input().split())
print(pBaseToQBase(num, a, b))