import math
import numpy as np
import networkx as nx
import random
import matplotlib.pyplot as plt
import matplotlib.colors as mcolors
from matplotlib import patches
import time

# Стандартное умножение
def mult(matrix1, matrix2):
    return np.dot(matrix1, matrix2)

# Логическое умножение
def mult_boolen(matrix1, matrix2):
    mult_matrix = mult(matrix1, matrix2)
    return np.vectorize(lambda x: True if x else False)(mult_matrix)

# Тропическое умножение
def mult_tropic(matrix1: np.ndarray, matrix2: np.ndarray):
    ans = []
    for i in range(len(matrix1)):
        for j in range(len(matrix2[0])):
            ans.append(min(matrix1[i] + matrix2[:, j]))
    return np.array(ans).reshape((len(matrix1), len(matrix2[0])))

def tropical_multiply(A, B):
    n = len(A)
    C = [[math.inf]*n for _ in range(n)]
    for i in range(n):
        for j in range(n):
            for k in range(n):
                C[i][j] = min(C[i][j], A[i][k] + B[k][j])
    return C

def visualize_tropical_matrix(matrix, ax, title="Тропическая матрица", cmap="plasma"):
    finite_vals = [val for row in matrix for val in row if not math.isinf(val)]
    max_val = max(finite_vals) if finite_vals else 1
    processed = [[max_val + 1 if math.isinf(val) else val for val in row] for row in matrix]
    text_matrix = [["∞" if math.isinf(val) else str(int(val)) for val in row] for row in matrix]

    arr = np.array(processed)
    thresh = (arr.min() + arr.max()) / 2
    n = arr.shape[0]

    im = ax.imshow(arr, cmap=cmap)
    ax.set_title(title)

    ax.set_xticks(np.arange(n))
    ax.set_yticks(np.arange(n))
    ax.set_xticklabels([])
    ax.set_yticklabels([])

    ax.set_xlabel("")
    ax.set_ylabel("")

    for i in range(n):
        for j in range(n):
            color = "black" if arr[i,j] >= thresh else "white"
            ax.text(j, i, text_matrix[i][j], ha="center", va="center", color=color)

    return im

def plot_matrix(matrix, title, cmap=None, binary=False, custom_cmap=None):
    plt.figure(figsize=(8, 8))
    ax = plt.gca()
    
    if cmap is None and custom_cmap is not None:
        cmap = custom_cmap
    
    im = ax.imshow(matrix, cmap=cmap, interpolation="nearest")
    
    for i in range(matrix.shape[0]+1):
        ax.axhline(i-0.5, color='black', linewidth=1)
        ax.axvline(i-0.5, color='black', linewidth=1)
    
    plt.title(title)
    ax.set_xticks([])
    ax.set_yticks([])
    
    for i in range(matrix.shape[0]):
        for j in range(matrix.shape[1]):
            value = matrix[i, j]
            text = "∞" if value == np.inf else ("True" if value and binary else ("False" if not value and binary else str(int(value))))
            ax.text(j, i, text, ha="center", va="center", color="black")

    plt.show()

def create_custom_cmap():
    colors = [(0.9, 0.9, 0.9), (0.5, 0.5, 0.5)]  # Серый для 0, синий для 1
    n_bins = 100
    cmap_name = 'gray_blue'
    cm = mcolors.LinearSegmentedColormap.from_list(cmap_name, colors, N=n_bins)
    return cm

def tropical_closure_with_power(adj_matrix):
    matrix_trop = [[math.inf if x == 0 else x for x in row] for row in adj_matrix]
    for i in range(len(matrix_trop)):
        matrix_trop[i][i] = 0
    
    matrix_power = [row[:] for row in matrix_trop]
    previous_matrix = None
    power = 1

    matrices = []
    custom_cmap = create_custom_cmap()  # Создаем пользовательскую цветовую схему

    while True:
        previous_matrix = [row[:] for row in matrix_power]
        matrix_power = tropical_multiply(matrix_power, matrix_trop)
        power += 1
        
        if previous_matrix == matrix_power:
            break
            
        matrices.append(([row[:] for row in matrix_power], power))
    
    # Визуализация всех шагов с новой цветовой схемой
    cols = 2
    rows = math.ceil(len(matrices) / cols)
    fig, axes = plt.subplots(rows, cols, figsize=(cols*7, rows*6))

    for idx, (matrix, pwr) in enumerate(matrices):
        ax = axes.flat[idx]
        
        # Преобразуем матрицу для отображения (заменяем inf на 0)
        display_matrix = np.array([[0 if math.isinf(x) else x for x in row] for row in matrix])
        
        # Находим максимальное конечное значение для нормализации цветов
        max_val = max([x for row in matrix for x in row if not math.isinf(x)] or [1])
        
        im = ax.imshow(display_matrix, cmap=custom_cmap, vmin=0, vmax=max_val)
        
        # Добавляем текст
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                value = matrix[i][j]
                text = "∞" if math.isinf(value) else str(int(value))
                ax.text(j, i, text, ha="center", va="center", color="black")
        
        ax.set_title(f"Степень {pwr}")
        ax.set_xticks([])
        ax.set_yticks([])

    for idx in range(len(matrices), rows*cols):
        fig.delaxes(axes.flat[idx])

    plt.tight_layout()
    plt.show()

    print(f"Матрица стабилизировалась на степени: {power-1}")
    return matrix_power

def floyd_warshall(adj_matrix):
    n = len(adj_matrix)
    dist = [[math.inf] * n for _ in range(n)]
    
    for i in range(n):
        for j in range(n):
            if adj_matrix[i][j] != 0:
                dist[i][j] = adj_matrix[i][j]
        dist[i][i] = 0
    
    for k in range(n):
        for i in range(n):
            for j in range(n):
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
    
    return dist

def compare_algorithms_time(adj_matrix):
    # Измеряем время работы алгоритма Флойда-Уоршелла
    start_time_fw = time.time()
    fw_matrix = floyd_warshall(adj_matrix)
    end_time_fw = time.time()
    fw_time = end_time_fw - start_time_fw

    # Измеряем время работы тропического умножения
    start_time_tropic = time.time()
    closure_matrix = tropical_closure_with_power(adj_matrix)
    end_time_tropic = time.time()
    tropic_time = end_time_tropic - start_time_tropic

    print(f"Время работы алгоритма Флойда-Уоршелла: {fw_time:.6f} секунд")
    print(f"Время работы тропического умножения: {tropic_time:.6f} секунд")

    return fw_time, tropic_time

def heawood():
    G = nx.petersen_graph()
    D = nx.DiGraph()
    D.add_nodes_from(G.nodes())
    for u, v in G.edges():
        if random.choice([True, False]):  
            D.add_edge(u, v)
        else:
            D.add_edge(v, u)
    
    plt.figure(figsize=(8, 6))
    pos = nx.spring_layout(D)  
    nx.draw(D, pos, with_labels=True, node_color="gray", edge_color="black", arrows=True)
    plt.title("Граф Петерсона со случайной ориентацией рёбер")
    plt.show()

    adj_matrix = nx.to_numpy_array(D, dtype=int).tolist()
    
    custom_cmap = create_custom_cmap()
    plot_matrix(np.array(adj_matrix), "Матрица смежности графа Петерсона", custom_cmap=custom_cmap)

    k = int(input('Введите длину рассматриваемых путей: '))

    # Количество путей с той же цветовой схемой
    mult_matrix = np.array(adj_matrix)
    for _ in range(k-1):
        mult_matrix = mult(mult_matrix, adj_matrix)
    plot_matrix(mult_matrix, f"Количество путей при длине пути {k}", custom_cmap=custom_cmap)

    # Достижимость с той же цветовой схемой
    bool_matrix = np.array(adj_matrix)
    for _ in range(k-1):
        bool_matrix = mult(bool_matrix, adj_matrix)
    plot_matrix(bool_matrix, f"Достижимость до вершины при длине пути {k}", custom_cmap=custom_cmap, binary=True)

    # Минимальные пути (тропическое умножение)
    tropic_matrix = np.where(np.array(adj_matrix) == 0, np.inf, np.array(adj_matrix))  
    np.fill_diagonal(tropic_matrix, 0)
    tmp = tropic_matrix
    for _ in range(k - 1):
        tropic_matrix = mult_tropic(tropic_matrix, tmp)
    plot_matrix(tropic_matrix, f"Минимальный путь при длине не более {k}", custom_cmap=custom_cmap)

    # Тропическое замыкание с визуализацией степеней
    print("\nВычисление тропического замыкания:")
    closure_matrix = tropical_closure_with_power(adj_matrix)
    plot_matrix(np.array(closure_matrix), "Тропическое замыкание графа Петерсона", custom_cmap=custom_cmap)

    # Алгоритм Флойда-Уоршелла
    print("\nВычисление алгоритмом Флойда-Уоршелла:")
    fw_matrix = floyd_warshall(adj_matrix)
    plot_matrix(np.array(fw_matrix), "Матрица кратчайших путей (Флойд-Уоршелл)", custom_cmap=custom_cmap)

    # Сравнение результатов
    print("\nСравнение результатов:")
    closure_np = np.array(closure_matrix)
    fw_np = np.array(fw_matrix)
    
    # Проверка на равенство (с учётом inf и числовой погрешности)
    comparison = np.isclose(closure_np, fw_np, equal_nan=True)
    print("Матрицы совпадают:", np.all(comparison))
    
    # Визуализация различий (если есть)
    if not np.all(comparison):
        diff_matrix = np.where(comparison, 0, 1)
        plot_matrix(diff_matrix, "Различия между тропическим замыканием и Флойдом-Уоршеллом", binary=True)

    compare_algorithms_time(adj_matrix)

heawood()
