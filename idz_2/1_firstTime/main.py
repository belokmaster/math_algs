import networkx as nx
import random
import matplotlib.pyplot as plt
import numpy as np
from itertools import combinations

def kargers_algorithm(graph):
    """Одна итерация алгоритма Каргера"""
    contracted_graph = graph.copy()
    
    while contracted_graph.number_of_nodes() > 2:
        edges = list(contracted_graph.edges())
        if not edges:
            return 0
        u, v = random.choice(edges)
        contracted_graph = nx.contracted_edge(contracted_graph, (u, v), self_loops=False)
        contracted_graph.remove_edges_from(nx.selfloop_edges(contracted_graph))
    
    return contracted_graph.number_of_edges()

def exact_min_cut(graph):
    """Точный минимальный разрез"""
    return len(nx.minimum_edge_cut(graph))

def generate_erdos_renyi_graph(n, p):
    """Генератор графов Эрдёша-Реньи"""
    graph = nx.Graph()
    graph.add_nodes_from(range(n))
    
    for u, v in combinations(range(n), 2):
        if random.random() < p:
            graph.add_edge(u, v)
    
    return graph

def analyze_first_try_success(num_graphs=100, n=10, p=0.5):
    """Анализ вероятности успеха с первого запуска"""
    success_count = 0
    exact_min_cuts = []
    first_try_results = []
    
    for _ in range(num_graphs):
        graph = generate_erdos_renyi_graph(n, p)
        
        if not nx.is_connected(graph):
            continue
            
        exact_cut = exact_min_cut(graph)
        first_try_cut = kargers_algorithm(graph)
        
        exact_min_cuts.append(exact_cut)
        first_try_results.append(first_try_cut)
        
        if first_try_cut == exact_cut:
            success_count += 1
    
    success_prob = success_count / num_graphs
    print(f"Вероятность успеха с первого раза: {success_prob:.4f}")
    print(f"Теоретическая оценка:  {2/(n**2):.4f}")  # Исправлено на 2/n²
    
    # Создаем один график для сравнения
    plt.figure(figsize=(10, 6))
    
    # График сравнения результатов
    plt.scatter(exact_min_cuts, first_try_results, alpha=0.7, 
               label='Результаты первого запуска')
    plt.plot([min(exact_min_cuts), max(exact_min_cuts)], 
             [min(exact_min_cuts), max(exact_min_cuts)], 
             'r--', label='Идеальное совпадение')
    
    # Подсчет и отображение случаев успеха
    successes = [(x, y) for x, y in zip(exact_min_cuts, first_try_results) if x == y]
    if successes:
        xs, ys = zip(*successes)
        plt.scatter(xs, ys, color='green', marker='o', s=100,
                   label=f'Успешные попытки ({len(successes)}/{num_graphs})',
                   edgecolors='black')
    
    plt.xlabel("Точный минимальный разрез", fontsize=12)
    plt.ylabel("Результат первого запуска Каргера", fontsize=12)
    plt.title(f"Сравнение первого запуска алгоритма Каргера с точным решением\n(n={n}, p={p}, успехов: {success_prob:.2%})", fontsize=14)
    plt.legend(fontsize=10)
    plt.grid(True, alpha=0.3)
    
    plt.tight_layout()
    plt.show()

# Запуск анализа
print("Анализ вероятности успеха с первого запуска (n=20, p=0.5)")
analyze_first_try_success(num_graphs=2000, n=20, p=0.5)