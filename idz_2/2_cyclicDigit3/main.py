import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
import math

# Функция для проверки связности и вычисления M
def analyze_graph(n, p, trials):
    connected_count = 0
    M_values = []
    for _ in range(trials):
        G = nx.erdos_renyi_graph(n, p)
        if nx.is_connected(G):
            connected_count += 1
        m = G.number_of_edges()
        k = nx.number_connected_components(G)
        M = m - n + k
        M_values.append(M)
    connectivity_prob = connected_count / trials
    avg_M = sum(M_values) / trials
    return connectivity_prob, avg_M

# Параметры
n_values = [50, 100, 200]
trials = 100
p_range = np.linspace(0, 0.1, 50)

# Словари для результатов
connectivity_results = {n: [] for n in n_values}
M_results = {n: [] for n in n_values}

# Генерация данных
for n in n_values:
    print(f"Обрабатываем n = {n}")
    for p in p_range:
        conn_prob, avg_M = analyze_graph(n, p, trials)
        connectivity_results[n].append(conn_prob)
        M_results[n].append(avg_M)

# Построение графиков
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(14, 6))

# График 1: Доля связных графов
for n in n_values:
    ax1.plot(p_range, connectivity_results[n], label=f'n = {n}')
    p_threshold = math.log(n) / n
    ax1.axvline(x=p_threshold, color='gray', linestyle='--', 
                label=f'ln({n})/{n} ≈ {p_threshold:.3f}' if n == n_values[0] else None)
ax1.set_xlabel('Вероятность p')
ax1.set_ylabel('Доля связных графов')
ax1.set_title('Вероятность связности графа')
ax1.legend()
ax1.grid()

# График 2: Среднее M
for n in n_values:
    ax2.plot(p_range, M_results[n], label=f'n = {n}')
    p_threshold = math.log(n) / n
    ax2.axvline(x=p_threshold, color='gray', linestyle='--', 
                label=f'ln({n})/{n} ≈ {p_threshold:.3f}' if n == n_values[0] else None)
ax2.set_xlabel('Вероятность p')
ax2.set_ylabel('Среднее цикломатическое число ')
ax2.set_title('Зависимость μ от p')
ax2.legend()
ax2.grid()

plt.tight_layout()
plt.show()