import networkx as nx
import random
import matplotlib.pyplot as plt
import math
from itertools import combinations

def kargers_algorithm(graph):
    """
    Реализация алгоритма Каргера для нахождения минимального разреза
    """
    # Создаем копию графа, чтобы не изменять исходный
    contracted_graph = graph.copy()
    
    # Пока в графе больше 2 вершин
    while contracted_graph.number_of_nodes() > 2:
        # Выбираем случайное ребро
        edges = list(contracted_graph.edges())
        if not edges:
            return 0  # Граф несвязный
        u, v = random.choice(edges)
        
        # Контрактируем ребро (объединяем вершины u и v)
        contracted_graph = nx.contracted_edge(contracted_graph, (u, v), self_loops=False)
        
        # Удаляем self-loops (петли)
        contracted_graph.remove_edges_from(nx.selfloop_edges(contracted_graph))
    
    # Возвращаем количество оставшихся ребер (разрез)
    return contracted_graph.number_of_edges()

def kargers_algorithm_repeated(graph, repetitions=None):
    """
    Многократный запуск алгоритма Каргера для увеличения вероятности нахождения минимального разреза
    """
    min_cut = float('inf')
    n = graph.number_of_nodes()
    
    if repetitions is None:
        # Согласно гипотезе, вероятность успеха на одной итерации >= 1/n^2
        # Поэтому нужно порядка n^2 log n повторений для высокой вероятности успеха
        repetitions = int(n**2 * math.log(n)) + 1
    
    for _ in range(repetitions):
        current_cut = kargers_algorithm(graph)
        if current_cut < min_cut:
            min_cut = current_cut
            if min_cut == 0:  # Граф несвязный
                return 0
    
    return min_cut

def generate_erdos_renyi_graph(n, p):
    """
    Генерация случайного графа модели Эрдёша-Реньи G(n, p)
    """
    graph = nx.Graph()
    graph.add_nodes_from(range(n))
    
    # Добавляем ребра с вероятностью p
    for u, v in combinations(range(n), 2):
        if random.random() < p:
            graph.add_edge(u, v)
    
    return graph

def visualize_graph_with_cut(graph, cut_edges=None):
    """
    Визуализация графа с выделением разреза (если задан)
    """
    pos = nx.spring_layout(graph)
    plt.figure(figsize=(10, 8))
    
    # Рисуем все ребра
    nx.draw_networkx_edges(graph, pos, alpha=0.3)
    
    # Если заданы ребра разреза, рисуем их красным
    if cut_edges:
        nx.draw_networkx_edges(graph, pos, edgelist=cut_edges, edge_color='r', width=2)
    
    # Рисуем вершины
    nx.draw_networkx_nodes(graph, pos, node_size=200, node_color='skyblue')
    nx.draw_networkx_labels(graph, pos)
    
    plt.title("Граф с минимальным разрезом (красные ребра)")
    plt.axis('off')
    plt.show()

def find_and_visualize_min_cut(n, p):
    """
    Генерация графа, поиск минимального разреза и визуализация
    """
    # Генерируем граф
    graph = generate_erdos_renyi_graph(n, p)
    
    # Проверяем связность
    if not nx.is_connected(graph):
        print("Граф несвязный, минимальный разрез = 0")
        visualize_graph_with_cut(graph)
        return
    
    # Находим минимальный разрез
    min_cut = kargers_algorithm_repeated(graph)
    print(f"Найден минимальный разрез размера {min_cut}")
    
    # Для визуализации найдем один из таких разрезов (может быть не минимальным)
    # Это демонстрационная часть - в реальности алгоритм Каргера не сохраняет разрез
    cut_edges = []
    if min_cut > 0:
        # Запустим алгоритм еще раз, чтобы получить ребра разреза
        while len(cut_edges) != min_cut:
            cut_edges = []
            contracted_graph = graph.copy()
            while contracted_graph.number_of_nodes() > 2:
                edges = list(contracted_graph.edges())
                u, v = random.choice(edges)
                contracted_graph = nx.contracted_edge(contracted_graph, (u, v), self_loops=False)
                contracted_graph.remove_edges_from(nx.selfloop_edges(contracted_graph))
            cut_edges = list(contracted_graph.edges())
    
    # Визуализируем граф с разрезом
    visualize_graph_with_cut(graph, cut_edges)

# Пример использования
n = 15  # Количество вершин
p = 0.3  # Вероятность образования ребра

find_and_visualize_min_cut(n, p)