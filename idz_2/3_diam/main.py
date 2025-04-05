import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
import math

# Функция для вычисления вероятности связности и среднего диаметра
def analyze_graph(n, p, trials):
    connected_count = 0
    diameters = []
    for _ in range(trials):
        G = nx.erdos_renyi_graph(n, p)
        if nx.is_connected(G):  # Проверяем связность
            connected_count += 1
            diam = nx.diameter(G)  # Считаем диаметр для связного графа
            diameters.append(diam)
    connectivity_prob = connected_count / trials
    avg_diameter = sum(diameters) / connected_count if connected_count > 0 else 0
    return connectivity_prob, avg_diameter

# Параметры для первого эксперимента
n_values_small = [50, 75, 100, 125]
p_range = np.arange(0, 1.01, 0.05)
trials = 100

# Словари для результатов
connectivity_results_small = {n: [] for n in n_values_small}
diameter_results_small = {n: [] for n in n_values_small}

# Первый эксперимент (n = 5, 10, 15, 20)
for n in n_values_small:
    print(f"Обрабатываем n = {n} (малые значения)")
    for p in p_range:
        conn_prob, avg_diam = analyze_graph(n, p, trials)
        connectivity_results_small[n].append(conn_prob)
        diameter_results_small[n].append(avg_diam)

# Построение графиков
fig, (ax1, ax2) = plt.subplots(1, 2, figsize=(14, 6))  # Убираем 2 пустых квадрата

# График 1: Вероятность связности (малые n)
for n in n_values_small:
    ax1.plot(p_range, connectivity_results_small[n], label=f'n = {n}')
    p_cycle = 1 / n
    ax1.axvline(x=p_cycle, color='red', linestyle='--', 
                label=f'1/{n} ≈ {p_cycle:.2f}' if n == n_values_small[0] else None)
    p_conn = math.log(n) / n
    ax1.axvline(x=p_conn, color='blue', linestyle='--', 
                label=f'ln({n})/{n} ≈ {p_conn:.2f}' if n == n_values_small[0] else None)
ax1.set_xlabel('Вероятность p')
ax1.set_ylabel('Доля связных графов')
ax1.set_title('Вероятность связности (n = 5, 10, 15, 20)')
ax1.legend()
ax1.grid()

# График 2: Средний диаметр (малые n)
for n in n_values_small:
    ax2.plot(p_range, diameter_results_small[n], label=f'n = {n}')
    p_cycle = 1 / n
    ax2.axvline(x=p_cycle, color='red', linestyle='--')
    p_conn = math.log(n) / n
    ax2.axvline(x=p_conn, color='blue', linestyle='--')
ax2.set_xlabel('Вероятность p')
ax2.set_ylabel('Средний диаметр')
ax2.set_title('Диаметр связных графов (n = 5, 10, 15, 20)')
ax2.legend()
ax2.grid()

plt.tight_layout()
plt.show()
