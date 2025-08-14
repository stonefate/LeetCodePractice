# LeetCode 练习项目

使用Go 1.24.5和Python 3.13进行LeetCode算法题目练习。

## 项目结构

```
LeetCodePractice/
├── go/                    # Go语言解答
│   ├── easy/             # 简单题目
│   ├── medium/           # 中等题目  
│   ├── hard/             # 困难题目
│   ├── utils/            # 通用工具函数
│   ├── tests/            # 测试文件
│   └── go.mod            # Go模块文件
├── python/               # Python语言解答
│   ├── easy/             # 简单题目
│   ├── medium/           # 中等题目
│   ├── hard/             # 困难题目
│   ├── utils/            # 通用工具函数
│   ├── tests/            # 测试文件
│   └── pyproject.toml    # Python项目配置 (uv管理)
├── docs/                 # 解题思路文档
│   ├── easy/
│   ├── medium/
│   └── hard/
├── CLAUDE.md            # Claude助手项目指南
└── README.md            # 项目说明
```

## 快速开始

### Go环境
```bash
cd go
go mod tidy
go test ./...
```

### Python环境 (使用uv)
```bash
cd python
uv sync
uv run pytest
```

## 开发工作流

1. 选择题目并在对应难度目录创建解答文件
2. 同时实现Go和Python版本
3. 编写测试用例
4. 记录解题思路到docs目录
5. 提交代码

详细指南请参考 [CLAUDE.md](./CLAUDE.md)