import re
from math import sin, cos, log as ln_func, exp

# --- ДОБАВЛЕНИЕ ДЛЯ ВИЗУАЛИЗАЦИИ ---
GRAPHVIZ_AVAILABLE = False
try:
    import graphviz
    GRAPHVIZ_AVAILABLE = True
except ImportError:
    print("Библиотека graphviz не найдена. Визуализация AST в PNG будет недоступна.")
    print("Пожалуйста, установите ее: pip install graphviz")
    print("Также убедитесь, что сама программа Graphviz установлена в системе и добавлена в PATH.")

# Глобальный счетчик для уникальных имен узлов в DOT, чтобы они не пересекались в одном графе
dot_node_counter_for_graph = 0

def reset_dot_node_counter_for_graph():
    global dot_node_counter_for_graph
    dot_node_counter_for_graph = 0

def get_unique_dot_node_name_for_dot():
    global dot_node_counter_for_graph
    dot_node_counter_for_graph += 1
    return f"n{dot_node_counter_for_graph}"

def ast_to_dot_recursive(current_ast_node, dot_graph_obj, parent_dot_name=None, edge_label=""):
    current_dot_name = get_unique_dot_node_name_for_dot()
    node_display_value = current_ast_node.value
    if current_ast_node.type == 'const':
        if isinstance(current_ast_node.value, float):
            if current_ast_node.value.is_integer():
                node_display_value = int(current_ast_node.value)
            else:
                node_display_value = f"{current_ast_node.value:.3g}"
                if node_display_value.endswith(".0"):
                     node_display_value = node_display_value[:-2]

    label_text = f"[{current_ast_node.type}]\n{node_display_value}"
    dot_graph_obj.node(current_dot_name, label=label_text, shape="box", style="rounded,filled", fillcolor="lightblue")

    if parent_dot_name:
        dot_graph_obj.edge(parent_dot_name, current_dot_name, label=edge_label)

    if len(current_ast_node.children) == 1:
        ast_to_dot_recursive(current_ast_node.children[0], dot_graph_obj, current_dot_name, "Arg")
    elif len(current_ast_node.children) == 2:
        ast_to_dot_recursive(current_ast_node.children[0], dot_graph_obj, current_dot_name, "L")
        ast_to_dot_recursive(current_ast_node.children[1], dot_graph_obj, current_dot_name, "R")

def save_ast_to_png(ast_root_node, filename_prefix):
    if not GRAPHVIZ_AVAILABLE: # Проверяем флаг доступности библиотеки
        # print(f"Graphviz библиотека не доступна, пропуск сохранения {filename_prefix}.png")
        return

    if ast_root_node is None:
        # print(f"AST узел пуст, пропуск сохранения {filename_prefix}.png")
        return

    reset_dot_node_counter_for_graph()
    # graphviz здесь уже импортирован, если GRAPHVIZ_AVAILABLE is True
    dot_diagram = graphviz.Digraph(comment=f'AST for {filename_prefix}')
    dot_diagram.attr(rankdir='TB', labelloc='t', label=f'AST: {filename_prefix.split("_")[-1]}')

    ast_to_dot_recursive(ast_root_node, dot_diagram)

    try:
        dot_diagram.render(filename_prefix, cleanup=True, format='png', view=False)
        print(f"AST сохранено в {filename_prefix}.png")
    except graphviz.backend.execute.ExecutableNotFound:
        print("\n!!! Ошибка: Утилита 'dot' (часть Graphviz) не найдена в системном PATH. !!!")
        print("Пожалуйста, убедитесь, что Graphviz корректно установлен и его директория bin добавлена в PATH.")
        print(f"Визуализация для {filename_prefix}.png будет пропущена.\n")
        # Не меняем GRAPHVIZ_AVAILABLE здесь, чтобы не усложнять.
        # Ошибка будет повторяться для каждого файла, если 'dot' не найден.
    except Exception as e:
        print(f"Непредвиденная ошибка при рендеринге AST в PNG ({filename_prefix}): {e}")
        import traceback
        traceback.print_exc()


# --- Структура узла AST ---
class Node:
    def __init__(self, type, value, children=None):
        self.type = type
        self.value = value
        self.children = children if children is not None else []

    def __repr__(self, level=0, prefix="Root:"):
        ret = "\t" * level + prefix + f" [{self.type}] {self.value}\n"
        for i, child in enumerate(self.children):
            child_prefix = "L:" if len(self.children) > 1 and i == 0 else \
                           "R:" if len(self.children) > 1 and i == 1 else \
                           "Arg:"
            ret += child.__repr__(level + 1, prefix=child_prefix)
        return ret

    def copy(self):
        return Node(self.type, self.value, [child.copy() for child in self.children])

# --- Парсер (Инфикс -> AST) ---
OPERATORS = {
    '+': {'prec': 1, 'assoc': 'L', 'type': 'binary'},
    '-': {'prec': 1, 'assoc': 'L', 'type': 'binary_or_unary'},
    '*': {'prec': 2, 'assoc': 'L', 'type': 'binary'},
    '/': {'prec': 2, 'assoc': 'L', 'type': 'binary'},
    '^': {'prec': 3, 'assoc': 'R', 'type': 'binary'},
    'sin': {'prec': 4, 'type': 'unary_func'},
    'cos': {'prec': 4, 'type': 'unary_func'},
    'ln': {'prec': 4, 'type': 'unary_func'},
    'exp': {'prec': 4, 'type': 'unary_func'},
}
UNARY_FUNCTIONS = {'sin', 'cos', 'ln', 'exp'}

class Parser:
    def __init__(self, expression_string):
        self.tokens = self._tokenize(expression_string)
        self.pos = 0

    def _tokenize(self, expression_string):
        token_specification = [
            ('NUMBER',   r'\d+(\.\d*)?'),
            ('VARIABLE', r'[a-zA-Z_][a-zA-Z0-9_]*'),
            ('OP',       r'[+\-*/^()]'),
            ('WHITESPACE',r'\s+'),
            ('MISMATCH', r'.'),
        ]
        tok_regex = '|'.join('(?P<%s>%s)' % pair for pair in token_specification)
        tokens = []
        for mo in re.finditer(tok_regex, expression_string):
            kind = mo.lastgroup
            value = mo.group()
            if kind == 'NUMBER':
                tokens.append({'type': 'CONST', 'value': float(value)})
            elif kind == 'VARIABLE':
                if value in UNARY_FUNCTIONS:
                    tokens.append({'type': 'FUNC', 'value': value})
                else:
                    tokens.append({'type': 'VAR', 'value': value})
            elif kind == 'OP':
                tokens.append({'type': 'OP', 'value': value})
            elif kind == 'WHITESPACE':
                continue
            elif kind == 'MISMATCH':
                raise ValueError(f"Неожиданный символ: '{value}'")
        return tokens

    def _peek(self):
        return self.tokens[self.pos] if self.pos < len(self.tokens) else None

    def _consume(self, expected_type=None, expected_value=None):
        if self.pos >= len(self.tokens):
            expected_str = ""
            if expected_type: expected_str += f"тип {expected_type}"
            if expected_value: expected_str += f" значение '{expected_value}'"
            raise ValueError(f"Неожиданный конец выражения, ожидался {expected_str}")

        token = self.tokens[self.pos]
        if expected_type and token['type'] != expected_type:
            raise ValueError(f"Ожидался тип {expected_type}, получен {token['type']} ('{token['value']}')")
        if expected_value and token['value'] != expected_value:
            raise ValueError(f"Ожидалось значение '{expected_value}', получено '{token['value']}'")
        self.pos += 1
        return token

    def parse(self):
        if not self.tokens:
            return Node('const', 0)
        ast = self._parse_add_sub()
        if self.pos < len(self.tokens):
            raise ValueError(f"Лишние токены в конце выражения, начиная с: {self.tokens[self.pos]}")
        return ast

    def _parse_add_sub(self):
        node = self._parse_mul_div()
        while True:
            token = self._peek()
            if token and token['type'] == 'OP' and token['value'] in ('+', '-'):
                op_val = token['value']
                self._consume('OP', op_val)
                right_node = self._parse_mul_div()
                node = Node('op', op_val, [node, right_node])
            else:
                break
        return node

    def _parse_mul_div(self):
        node = self._parse_power()
        while True:
            token = self._peek()
            if token and token['type'] == 'OP' and token['value'] in ('*', '/'):
                op_val = token['value']
                self._consume('OP', op_val)
                right_node = self._parse_power()
                node = Node('op', op_val, [node, right_node])
            else:
                break
        return node

    def _parse_power(self):
        node = self._parse_unary()
        token = self._peek()
        if token and token['type'] == 'OP' and token['value'] == '^':
            op_val = token['value']
            self._consume('OP', op_val)
            right_node = self._parse_power()
            node = Node('op', op_val, [node, right_node])
        return node

    def _parse_unary(self):
        token = self._peek()
        if token is None:
             raise ValueError("Неожиданный конец выражения! Ожидался унарный оператор!")

        if token['type'] == 'OP' and token['value'] == '-':
            self._consume('OP', '-')
            operand_node = self._parse_unary()
            return Node('op', 'neg', [operand_node])
        elif token['type'] == 'FUNC':
            func_name = token['value']
            self._consume('FUNC')
            if not (self._peek() and self._peek()['value'] == '('):
                raise ValueError(f"Ожидалась '(' после функции {func_name}")
            self._consume('OP', '(')
            arg_node = self._parse_add_sub()
            if not (self._peek() and self._peek()['value'] == ')'):
                raise ValueError(f"Ожидалась ')' после аргумента функции {func_name}")
            self._consume('OP', ')')
            return Node('op', func_name, [arg_node])
        else:
            return self._parse_primary_atom()

    def _parse_primary_atom(self):
        token = self._peek()
        if token is None:
            raise ValueError("Неожиданный конец выражения, ожидалось атомарное выражение.")

        if token['type'] == 'CONST':
            self._consume('CONST')
            return Node('const', token['value'])
        elif token['type'] == 'VAR':
            self._consume('VAR')
            return Node('var', token['value'])
        elif token['type'] == 'OP' and token['value'] == '(':
            self._consume('OP', '(')
            expr_node = self._parse_add_sub()
            if not (self._peek() and self._peek()['value'] == ')'):
                raise ValueError("Ожидалась закрывающая скобка ')'")
            self._consume('OP', ')')
            return expr_node
        else:
            raise ValueError(f"Неожиданный токен при парсинге атом. выражения: {token}")

# --- Преобразование AST -> Инфикс ---
PRINT_PRECEDENCE = {
    '+': 1, '-': 1,
    'neg': 2,
    '*': 2, '/': 2,
    '^': 3,
    'sin': 4, 'cos': 4, 'ln': 4, 'exp': 4
}

def tree_to_infix(node, parent_op_prec=0, is_left_child_of_parent=False):
    if node.type == 'const':
        val = node.value
        if isinstance(val, float):
            if val.is_integer(): return str(int(val))
            return f"{val:.4g}".rstrip('0').rstrip('.')
        return str(val)
    elif node.type == 'var':
        return node.value
    elif node.type == 'op':
        op = node.value
        current_op_prec = PRINT_PRECEDENCE.get(op, 99)

        if op in UNARY_FUNCTIONS:
            arg_str = tree_to_infix(node.children[0], current_op_prec)
            return f"{op}({arg_str})"
        elif op == 'neg':
            arg_node = node.children[0]
            arg_str = tree_to_infix(arg_node, current_op_prec)
            if arg_node.type == 'op' and len(arg_node.children) == 2 and \
               PRINT_PRECEDENCE.get(arg_node.value, 99) < current_op_prec:
                return f"-({arg_str})"
            return f"-{arg_str}"
        else: # Бинарные операции
            left_str = tree_to_infix(node.children[0], current_op_prec, True)
            right_str = tree_to_infix(node.children[1], current_op_prec, False)
            expr_str = f"{left_str} {op} {right_str}"
            needs_paren = False
            if current_op_prec < parent_op_prec:
                needs_paren = True
            elif current_op_prec == parent_op_prec:
                current_op_info = OPERATORS.get(op)
                if current_op_info:
                    if current_op_info['assoc'] == 'L' and not is_left_child_of_parent:
                        needs_paren = True
                    elif current_op_info['assoc'] == 'R' and is_left_child_of_parent:
                        needs_paren = True
                if op in ('-', '/') and not is_left_child_of_parent:
                     needs_paren = True
            return f"({expr_str})" if needs_paren else expr_str
    return "ERROR_UNKNOWN_NODE_TYPE"

def u_original_copy(node):
    return node.copy()

def d_const(node, var_name):
    return Node('const', 0)

def d_var(node, var_name):
    return Node('const', 1) if node.value == var_name else Node('const', 0)

def d_op(node, var_name):
    op = node.value
    children = node.children
    u = children[0]
    v = children[1] if len(children) > 1 else None
    du = differentiate(u, var_name)
    dv = differentiate(v, var_name) if v else None

    if op == '+': return Node('op', '+', [du, dv])
    if op == '-': return Node('op', '-', [du, dv])
    if op == 'neg': return Node('op', 'neg', [du])
    if op == '*':
        return Node('op', '+', [Node('op', '*', [du, u_original_copy(v)]),
                                Node('op', '*', [u_original_copy(u), dv])])
    if op == '/':
        num = Node('op', '-', [Node('op', '*', [du, u_original_copy(v)]),
                               Node('op', '*', [u_original_copy(u), dv])])
        den = Node('op', '^', [u_original_copy(v), Node('const', 2)])
        return Node('op', '/', [num, den])
    if op == '^':
        if v.type == 'const':
            c_val = v.value
            if c_val == 0: return Node('const', 0)
            if c_val == 1: return du
            term_pow = Node('op', '^', [u_original_copy(u), Node('const', c_val - 1)])
            term_coeff = Node('op', '*', [Node('const', c_val), term_pow])
            return Node('op', '*', [term_coeff, du])
        elif u.type == 'const':
            c_val = u.value
            if c_val == 1: return Node('const', 0)
            term_exp = Node('op', '^', [u_original_copy(u), u_original_copy(v)])
            term_ln_c = Node('op', 'ln', [u_original_copy(u)])
            res_part = Node('op', '*', [term_exp, term_ln_c])
            return Node('op', '*', [res_part, dv])
        else:
            uv_term = Node('op', '^', [u_original_copy(u), u_original_copy(v)])
            ln_u_term = Node('op', 'ln', [u_original_copy(u)])
            term1_factor = Node('op', '*', [dv, ln_u_term])
            du_div_u_term = Node('op', '/', [du, u_original_copy(u)])
            term2_factor = Node('op', '*', [u_original_copy(v), du_div_u_term])
            sum_term = Node('op', '+', [term1_factor, term2_factor])
            return Node('op', '*', [uv_term, sum_term])
    if op == 'sin': return Node('op', '*', [Node('op', 'cos', [u_original_copy(u)]), du])
    if op == 'cos': return Node('op', '*', [Node('op', 'neg', [Node('op', 'sin', [u_original_copy(u)])]), du])
    if op == 'ln': return Node('op', '*', [Node('op', '/', [Node('const', 1), u_original_copy(u)]), du])
    if op == 'exp': return Node('op', '*', [Node('op', 'exp', [u_original_copy(u)]), du])
    raise ValueError(f"Неизвестная операция для дифференцирования: {op}")

def differentiate(node, var_name):
    if node.type == 'const': return d_const(node, var_name)
    if node.type == 'var': return d_var(node, var_name)
    if node.type == 'op': return d_op(node, var_name)
    raise TypeError(f"Неизвестный тип узла: {node.type}")

# --- Упрощение выражений ---
MAX_SIMPLIFICATION_PASSES = 10

def _are_structurally_equal(node1, node2):
    if node1 is node2: return True
    if node1.type != node2.type or node1.value != node2.value:
        return False
    if len(node1.children) != len(node2.children):
        return False
    for child1, child2 in zip(node1.children, node2.children):
        if not _are_structurally_equal(child1, child2):
            return False
    return True

def apply_simplification_rules(node):
    original_node = node
    if node.type != 'op':
        return node
    op = node.value
    children = node.children

    if all(child.type == 'const' for child in children):
        vals = [child.value for child in children]
        try:
            if op == '+': node = Node('const', vals[0] + vals[1])
            elif op == '-': node = Node('const', vals[0] - vals[1])
            elif op == '*': node = Node('const', vals[0] * vals[1])
            elif op == '/':
                if vals[1] == 0: pass
                else: node = Node('const', vals[0] / vals[1])
            elif op == '^':
                if vals[0] == 0 and vals[1] == 0: pass
                else: node = Node('const', vals[0] ** vals[1])
            elif op == 'neg': node = Node('const', -vals[0])
            elif op == 'sin': node = Node('const', sin(vals[0]))
            elif op == 'cos': node = Node('const', cos(vals[0]))
            elif op == 'ln':
                if vals[0] <= 0: pass
                else: node = Node('const', ln_func(vals[0]))
            elif op == 'exp': node = Node('const', exp(vals[0]))
        except (ValueError, OverflowError, ZeroDivisionError):
            pass
        if node is not original_node: return node

    if len(children) == 2:
        u, v = children[0], children[1]
        if op == '+':
            if u.type == 'const' and u.value == 0: return v
            if v.type == 'const' and v.value == 0: return u
        if op == '-':
            if v.type == 'const' and v.value == 0: return u
        if op == '*':
            if (u.type == 'const' and u.value == 0) or \
               (v.type == 'const' and v.value == 0):
                return Node('const', 0)
            if u.type == 'const' and u.value == 1: return v
            if v.type == 'const' and v.value == 1: return u
        if op == '/':
            if v.type == 'const' and v.value == 1: return u
            if u.type == 'const' and u.value == 0:
                if not (v.type == 'const' and v.value == 0):
                     return Node('const', 0)
        if op == '^':
            if v.type == 'const' and v.value == 0:
                if not (u.type == 'const' and u.value == 0):
                    return Node('const', 1)
            if v.type == 'const' and v.value == 1: return u
            if u.type == 'const' and u.value == 1: return Node('const', 1)
            if u.type == 'const' and u.value == 0:
                if v.type == 'const' and v.value > 0: return Node('const', 0)
        if op == '-' and _are_structurally_equal(u, v): return Node('const', 0)
        if op == '/':
            is_denominator_zero_const = (v.type == 'const' and v.value == 0)
            if not is_denominator_zero_const and u.type == 'op' and u.value == '*':
                if _are_structurally_equal(u.children[0], v): return u.children[1]
                if _are_structurally_equal(u.children[1], v): return u.children[0]
    elif len(children) == 1:
        u = children[0]
        if op == 'neg' and u.type == 'op' and u.value == 'neg':
            return u.children[0]
        if op == 'ln' and u.type == 'const' and u.value == 1: return Node('const', 0)
        if op == 'exp' and u.type == 'const' and u.value == 0: return Node('const', 1)
        if op == 'ln' and u.type == 'op' and u.value == 'exp': return u.children[0]
        if op == 'exp' and u.type == 'op' and u.value == 'ln': return u.children[0]
        if op == 'sin' and u.type == 'const' and u.value == 0: return Node('const', 0)
        if op == 'cos' and u.type == 'const' and u.value == 0: return Node('const', 1)
    return node

def simplify_single_pass(node_to_simplify):
    if node_to_simplify.type in ('const', 'var'):
        return node_to_simplify

    simplified_children_list = []
    children_changed_in_recursion = False
    for child_node in node_to_simplify.children:
        simplified_child = simplify_single_pass(child_node)
        simplified_children_list.append(simplified_child)
        if child_node is not simplified_child:
            children_changed_in_recursion = True

    current_processing_node = node_to_simplify
    if children_changed_in_recursion:
        current_processing_node = Node(node_to_simplify.type, node_to_simplify.value, simplified_children_list)

    node_after_rules = apply_simplification_rules(current_processing_node)

    if node_after_rules is not current_processing_node or children_changed_in_recursion :
        return node_after_rules
    else:
        return node_to_simplify

def simplify_iterative(node):
    current_node_state = node.copy()
    for i in range(MAX_SIMPLIFICATION_PASSES):
        new_node_state_after_pass = simplify_single_pass(current_node_state)
        if current_node_state is new_node_state_after_pass:
            break
        if _are_structurally_equal(current_node_state, new_node_state_after_pass) and i > 0 :
             current_node_state = new_node_state_after_pass
             break
        current_node_state = new_node_state_after_pass
    return current_node_state

# --- Сборка и Тестирование ---
def run_example(expression_str, var_to_diff='x', example_counter=0):
    print(f"\n--- Пример #{example_counter}: {expression_str} ---")
    safe_expr_part = re.sub(r'[^a-zA-Z0-9_]', '', expression_str)[:20]
    if not safe_expr_part: safe_expr_part = "expr"
    filename_base = f"out_example_{example_counter}_{safe_expr_part}"

    ast = None
    try:
        parser = Parser(expression_str)
        ast = parser.parse()
        print(f"Исходное (парсинг): {tree_to_infix(ast)}")
        save_ast_to_png(ast, f"{filename_base}_0_initial")
    except Exception as e:
        print(f"ОШИБКА ПАРСИНГА: {e}")
        import traceback; traceback.print_exc()
        return

    diff_ast = None
    if ast:
        try:
            print(f"\nДифференцируем по '{var_to_diff}':")
            diff_ast = differentiate(ast, var_to_diff)
            diff_infix_raw = tree_to_infix(diff_ast)
            print(f"Производная (до упрощения): {diff_infix_raw}")
            save_ast_to_png(diff_ast, f"{filename_base}_1_diff_raw")
        except Exception as e:
            print(f"ОШИБКА ДИФФЕРЕНЦИРОВАНИЯ: {e}")
            import traceback; traceback.print_exc()

    if diff_ast:
        try:
            print("\nУпрощаем производную:")
            simplified_diff_ast = simplify_iterative(diff_ast)
            simplified_diff_infix = tree_to_infix(simplified_diff_ast)
            print(f"Производная (после упрощения): {simplified_diff_infix}")
            save_ast_to_png(simplified_diff_ast, f"{filename_base}_2_diff_simplified")
        except Exception as e:
            print(f"ОШИБКА УПРОЩЕНИЯ: {e}")
            import traceback; traceback.print_exc()
    print("-" * 40)

if __name__ == '__main__':
    examples = [
        "2", "x", "y", "x + 2", "2 * x", "x * 2", "x + x", "x * x",
        "x ^ 2", "2 ^ x", "x / 2", "sin(x)", "cos(x)", "ln(x)", "exp(x)",
        "-x", "sin(2*x)", "x * sin(x)", "(x+1)*(x-1)", "ln(x^2)", "x^x",
        "1+x+x^2/2",
        "0 * x", "1 * x", "x + 0", "x - x",
        "sin(0*x)", "(2+3)*x", "-(x+1)", "--x", "ln(exp(x))", "exp(ln(x))",
        "1/x", "x/x"
    ]
    global_example_counter = 0
    for ex_str in examples:
        global_example_counter += 1
        run_example(ex_str, 'x', example_counter=global_example_counter)

    global_example_counter += 1
    run_example("y * x + y^2", "y", example_counter=global_example_counter)

    print("\n--- Сложные тесты парсера и приоритетов ---")
    global_example_counter += 1
    run_example("1+2*3-4/2^2", 'x', example_counter=global_example_counter)
    global_example_counter += 1
    run_example("sin(-x*cos(x^2))", 'x', example_counter=global_example_counter)
    global_example_counter += 1
    run_example("1 + -x * 2", 'x', example_counter=global_example_counter)
    global_example_counter += 1
    run_example("2^3^2", 'x', example_counter=global_example_counter)
    global_example_counter += 1
    run_example("(2^3)^2", 'x', example_counter=global_example_counter)

    print("\n--- Тесты только упрощения (без дифференцирования) ---")
    test_simplify_expr_str = "0*x + 1*(x+0) - (y-y) + (z^0)*sin(0*k) + ln(1) + exp(0) + --x + (2*3+4)/2 + (a*b)/b"
    global_example_counter +=1
    print(f"\n--- Пример #{global_example_counter} (только упрощение): {test_simplify_expr_str} ---")
    filename_base_simplify_test = f"out_example_{global_example_counter}_pure_simplify"
    try:
        parser_s = Parser(test_simplify_expr_str)
        ast_s = parser_s.parse()
        print(f"AST до упрощения (из строки '{test_simplify_expr_str}'):")
        print(f"Инфикс до упрощения: {tree_to_infix(ast_s)}")
        save_ast_to_png(ast_s, f"{filename_base_simplify_test}_0_before")

        simplified_ast_s = simplify_iterative(ast_s)
        print(f"\nAST после упрощения:")
        print(f"Инфикс после упрощения: {tree_to_infix(simplified_ast_s)}")
        save_ast_to_png(simplified_ast_s, f"{filename_base_simplify_test}_1_after")
    except Exception as e:
        print(f"ОШИБКА в тесте только упрощения: {e}")
        import traceback; traceback.print_exc()