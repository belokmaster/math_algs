import random
import networkx as nx
import numpy as np
import matplotlib.pyplot as plt
from networkx.algorithms.connectivity import stoer_wagner

def karger_min_cut(graph):
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
        if find(u) == find(v):
            edges.remove((u, v))
            continue
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
        return 0, set(graph.nodes()), set()
    rep1, rep2 = list(clusters_map.keys())[:2]
    cut_set1, cut_set2 = set(clusters_map[rep1]), set(clusters_map[rep2])
    cut_size = 0
    for u, v in graph.edges():
        if (u in cut_set1 and v in cut_set2) or (u in cut_set2 and v in cut_set1):
            cut_size += 1
    return cut_size, cut_set1, cut_set2

def karger_success_prob(n, p, trials=50, k=10):
    success_counts = np.zeros(k+1)
    graphs_count = 0
    for _ in range(trials):
        G = nx.gnp_random_graph(n, p)
        if G.number_of_edges() == 0 or not nx.is_connected(G):
            continue
        graphs_count += 1
        mincut_val, partition = stoer_wagner(G)
        mincut_val = int(mincut_val)
        results = [karger_min_cut(G)[0] for _ in range(k)]
        for j in range(1, k+1):
            if any(res == mincut_val for res in results[:j]):
                success_counts[j] += 1
    if graphs_count == 0:
        return [0] * (k+1)
    success_probs = success_counts / graphs_count
    return success_probs.tolist()

# Параметры эксперимента
n = 10
p = 0.5
trials = 50
k = 20

# Получаем данные
probs = karger_success_prob(n, p, trials, k)

# Строим график
plt.figure(figsize=(10, 6))
plt.plot(range(1, k+1), probs[1:], marker='o', color='gray', linewidth=2, markersize=8)

plt.xlabel("Количество запусков алгоритма", fontsize=12)
plt.ylabel("Вероятность нахождения минимального разреза", fontsize=12)
plt.title(f"Вероятность успеха алгоритма Каргера для n=50, p={p}", fontsize=14)
plt.grid(True, linestyle='--', alpha=0.7)
plt.xticks(range(1, k+1))
plt.ylim(0, 1.05)

# Добавляем значения точек на график
for i, prob in enumerate(probs[1:]):
    plt.text(i+1, prob+0.02, f"{prob:.2f}", ha='center', va='bottom', fontsize=10)

plt.tight_layout()
plt.show()