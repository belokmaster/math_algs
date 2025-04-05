import networkx as nx
import random
import matplotlib.pyplot as plt
import numpy as np
from itertools import combinations

def karger_min_cut(graph):
    """Реализация алгоритма Каргера с Union-Find (как в первом коде)"""
    parent = {v: v for v in graph.nodes()}
    rank = {v: 0 for v in graph.nodes()}
    
    def find(v):
        if parent[v] != v:
            parent[v] = find(parent[v])
        return parent[v]
    
    def union(u, v):
        ru, rv = find(u), find(v)
        if ru == rv:
            return False
        if rank[ru] < rank[rv]:
            parent[ru] = rv
        elif rank[ru] > rank[rv]:
            parent[rv] = ru
        else:
            parent[rv] = ru
            rank[ru] += 1
        return True

    edges = list(graph.edges())
    clusters = graph.number_of_nodes()
    
    while clusters > 2 and edges:
        u, v = random.choice(edges)
        if not union(u, v):
            continue
        clusters -= 1
        union(u, v)
        clusters -= 1
        new_edges = []
        for a, b in edges:
            if find(a) != find(b):
                new_edges.append((a, b))
        edges = new_edges

    clusters_map = {}
    for v in graph.nodes():
        rep = find(v)
        clusters_map.setdefault(rep, []).append(v)
    
    if len(clusters_map) < 2:
        return 0
    
    rep1, rep2 = list(clusters_map.keys())[:2]
    cut_set1, cut_set2 = set(clusters_map[rep1]), set(clusters_map[rep2])
    cut_size = 0
    for u, v in graph.edges():
        if (u in cut_set1 and v in cut_set2) or (u in cut_set2 and v in cut_set1):
            cut_size += 1
    return cut_size

def exact_min_cut(graph):
    """Точный минимальный разрез (аналог Stoer-Wagner из первого кода)"""
    return len(nx.minimum_edge_cut(graph))

def analyze_karger_success(n=10, p=0.5, trials=50, max_k=20):
    """Анализ вероятности успеха в зависимости от числа запусков"""
    success_counts = np.zeros(max_k + 1)
    valid_graphs = 0
    
    for _ in range(trials):
        graph = nx.gnp_random_graph(n, p)
        
        if not nx.is_connected(graph) or graph.number_of_edges() == 0:
            continue
            
        valid_graphs += 1
        true_mincut = exact_min_cut(graph)
        
        # Запускаем алгоритм Каргера до max_k раз и запоминаем успехи
        found = False
        for k in range(1, max_k + 1):
            if not found:
                cut = karger_min_cut(graph)
                if cut == true_mincut:
                    success_counts[k:] += 1
                    found = True
    
    if valid_graphs == 0:
        return [0] * (max_k + 1)
    
    success_probs = success_counts / valid_graphs
    return success_probs.tolist()

# Параметры эксперимента (как в первом коде)
n = 15
p = 0.8
trials = 2000
max_k = 20

# Получаем данные
probs = analyze_karger_success(n, p, trials, max_k)

# Строим график (как в первом коде)
plt.figure(figsize=(10, 6))
plt.plot(range(1, max_k + 1), probs[1:max_k + 1], marker='o', color='gray', linewidth=2, markersize=8)

plt.xlabel("Количество запусков алгоритма", fontsize=12)
plt.ylabel("Вероятность нахождения минимального разреза", fontsize=12)
plt.title(f"Вероятность успеха алгоритма Каргера для n={n}, p={p}", fontsize=14)
plt.grid(True, linestyle='--', alpha=0.7)
plt.xticks(range(1, max_k + 1))
plt.ylim(0, 1.05)

# Добавляем значения точек на график
for i, prob in enumerate(probs[1:max_k + 1]):
    plt.text(i + 1, prob + 0.02, f"{prob:.2f}", ha='center', va='bottom', fontsize=10)

plt.tight_layout()
plt.show()

# Дополнительно выводим вероятность успеха с первого раза
print(f"Вероятность успеха с первого запуска: {probs[1]:.2f}")
print(f"Теоретическая оценка (2/n²): {2/(n**2):.4f}")