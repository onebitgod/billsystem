# Billing System (CLI) in Go 
### For anyone Starting with Go

A simple **command-line billing system** written entirely in **Go** with **no external packages** — built using only core language features like functions, conditionals, loops, and basic I/O.

---

## 🚀 Features

- Add items with price and quantity to bill
- Calculate total bill
- Display detailed bill summary
- CLI-based user interaction

---

## 🧱 Built With

- Go (no external libraries or dependencies)
- Core Go features: `fmt`, `bufio`, `os`, basic control flow

---

## 🖥️ How It Works

The program runs in your terminal and allows you to:

1. Create Items (name, price)
2. Create Bill (items, tip, items quantity)
3. View, Edit and Delete Items
4. View, Edit and Delete Bills

---

## 📦 Getting Started

### Clone the repo and run

```bash
git clone https://github.com/yourusername/billsystem.git
cd billsystem
go run main.go
```

```bash
---------- Main Menu ----------

| b: Bills Menu | i: Items Menu | q: Exit |

Choose your option : 
```

### 1. Bills Menu 
b: Bills Menu

```bash
---------- Bill Menu ----------

| v: View Bills | n: New Bill | m: Main Menu |
```

v: View Bills

```bash
----------------------------------- Bills List ---------------------------------------
-------------------------------------------------------------------------------------
| S.No  | Customer Name  | BillID | Total items | CreatedAt          | Total Amount |
-------------------------------------------------------------------------------------
| 1     | OneBitGod       | 1      | 2           | 14 May 2025, 11:59 | ₹650         |
-------------------------------------------------------------------------------------
|                                                                    Total Bills: 1 |
-------------------------------------------------------------------------------------


Choose Below option or Enter Bill Id to proceed

| d: Delete Bill | b: Bill Menu | m: Main Menu | q: exit |

Choose your option : 
```

## CODE LIKE GOD
🤝 Contribute, Code, Learn


