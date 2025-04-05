import networkx as nx
import matplotlib.pyplot as plt
import numpy as np

# Функция для вычисления цикломатического числа
def cyclomatic_number(graph):
    # Формула для цикломатического числа: M = E - V + C
    # где E - количество рёбер, V - количество вершин, C - количество компонент связности
    return len(graph.edges()) - len(graph.nodes()) + nx.number_connected_components(graph)

# Параметры эксперимента
n = 20  # количество вершин
p_values = np.linspace(0, 1, 100)  # значения p от 0 до 1
cyclical_numbers = []  # список для хранения цикломатических чисел

# Генерация графов и подсчёт цикломатического числа для каждого p
for p in p_values:
    G = nx.gnp_random_graph(n, p)  # генерация случайного графа с p вероятностью
    cyclical_numbers.append(cyclomatic_number(G))  # вычисление цикломатического числа

# Визуализация результатов
plt.figure(figsize=(10, 6))
plt.plot(p_values, cyclical_numbers, label="Цикломатическое число", color='b', lw=2)
plt.xlabel("Вероятность p", fontsize=12)
plt.ylabel("Цикломатическое число", fontsize=12)
plt.title(f"Зависимость цикломатического числа от p для графа с n={n} вершинами", fontsize=14)
plt.grid(True)
plt.axvline(x=1/n, color='r', linestyle='--', label=f"p = 1/n ({1/n:.2f})")
plt.legend()
plt.tight_layout()
plt.show()
