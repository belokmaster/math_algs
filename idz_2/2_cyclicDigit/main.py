import networkx as nx
import numpy as np
import matplotlib.pyplot as plt

# Задаём значения n и p
n_values = [5, 10, 15, 20]
p_values = np.arange(0, 1.01, 0.05)  # от 0 до 1 с шагом 0.05
trials = 1000  # число экспериментов для каждой пары (n, p)

# Словарь для хранения результатов
results = {n: [] for n in n_values}

# Цикл по всем n
for n in n_values:
    print(f"Обрабатываем n = {n}")
    for p in p_values:
        M_sum = 0
        # Генерируем 1000 графов для текущих n и p
        for _ in range(trials):
            G = nx.erdos_renyi_graph(n, p)  # создаём случайный граф
            m = G.number_of_edges()  # число рёбер
            k = nx.number_connected_components(G)  # число компонент
            M = m - n + k  # цикломатическое число
            M_sum += M
        # Считаем среднее M
        E_M = M_sum / trials
        results[n].append(E_M)

# Построение графиков
plt.figure(figsize=(10, 6))  # размер графика
for n in n_values:
    plt.plot(p_values, results[n], label=f'n = {n}')
plt.xlabel('Вероятность p')
plt.ylabel('Среднее цикломатическое число')
plt.title('Зависимость μ от p для разных n')
plt.legend()  # добавляем легенду
plt.grid()  # сетка для удобства
plt.show()