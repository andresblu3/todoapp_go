<template>
    <v-container class="todo-app">
        <v-row justify="center">
            <v-col cols="12" md="8">
                <h1 class="text-center mb-4">Lista de Tareas</h1>
                <v-form @submit.prevent="addTodo" class="todo-form mb-4" v-model="formValid">
                    <v-text-field v-model="newTodo" :rules="inputRules" label="Agregar una nueva tarea" outlined dense
                        clearable append-inner-icon="mdi-plus" @click:append-inner="addTodo"></v-text-field>
                </v-form>
                <v-row>
                    <v-col cols="12" v-for="todo in todos" :key="todo.ID" class="mb-2">
                        <v-card class="todo-item animate__animated animate__fadeIn">
                            <v-card-text class="todo-content">
                                <v-checkbox v-model="todo.completed" @change="updateTodo(todo)" class="todo-checkbox"
                                    hide-details></v-checkbox>
                                <span :class="{ completed: todo.completed }" class="todo-text">{{ todo.title }}</span>
                                <v-btn icon @click="deleteTodoById(todo.ID)" class="todo-delete-btn">
                                    <v-icon color="primary">mdi-delete</v-icon>
                                </v-btn>
                            </v-card-text>
                        </v-card>
                    </v-col>
                </v-row>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
export default {
    data() {
        return {
            todos: [],
            newTodo: '',
            formValid: false,
            inputRules: [
                v => !!v || 'La tarea es obligatoria',
                v => (v && v.length <= 200) || 'La tarea debe tener menos de 200 caracteres',
                v => (v && v.trim().length > 0) || 'La tarea no puede contener solo espacios en blanco'
            ]
        }
    },
    methods: {
        async fetchTodos() {
            try {
                const response = await this.$http.get('/todos')
                this.todos = response.data
                console.log('Tareas obtenidas:', this.todos)
            } catch (error) {
                console.error('Error al obtener las tareas:', error)
            }
        },
        async addTodo() {
            if (this.newTodo.trim() === '') {
                return
            }
            try {
                const response = await this.$http.post('/todos', {
                    title: this.newTodo,
                    completed: false
                })
                this.todos.push(response.data)
                this.newTodo = ''
            } catch (error) {
                console.error('Error al agregar la tarea:', error)
            }
        },
        async updateTodo(todo) {
            try {
                await this.$http.put(`/todos/${todo.ID}`, todo)
            } catch (error) {
                console.error('Error al actualizar la tarea:', error)
            }
        },
        async deleteTodoById(id) {
            if (!id) {
                console.error('El ID es indefinido. No se puede eliminar la tarea.')
                return
            }
            try {
                const response = await this.$http.delete(`/todos/${id}`)
                console.log('Respuesta de eliminación:', response)
                this.todos = this.todos.filter(todo => todo.ID !== id)
                this.fetchTodos()
            } catch (error) {
                console.error('Error al eliminar la tarea:', error)
            }
        }
    },
    created() {
        this.fetchTodos()
    }
}
</script>

<style scoped>
body {
    background-color: #e0e7ef;
    font-family: 'Roboto', sans-serif;
}

.todo-app {
    max-width: 600px;
    margin: 50px auto;
    padding: 20px;
    border-radius: 12px;
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
    background-color: #ffffff;
}

h1 {
    font-family: 'Roboto', sans-serif;
    font-weight: bold;
    color: #2c3e50;
}

.v-card {
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s, transform 0.3s;
    background-color: #f5f7fa;
}

.v-card:hover {
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.2);
    transform: translateY(-4px);
}

.todo-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.todo-checkbox {
    margin-right: 16px;
}

.todo-text {
    flex: 1;
    font-size: 16px;
    color: #2c3e50;
}

.todo-delete-btn {
    transition: background-color 0.3s, color 0.3s;
}

.todo-delete-btn:hover {
    background-color: #1a2027;
    color: white;
}

.completed {
    text-decoration: line-through;
    color: #7f8c8d;
}

.v-text-field--outlined.v-input--dense .v-input__control {
    border-radius: 8px;
}

.v-text-field--outlined.v-input--dense .v-input__append-inner .v-icon {
    cursor: pointer;
    transition: color 0.3s;
}

.v-text-field--outlined.v-input--dense .v-input__append-inner .v-icon:hover {
    color: #2980b9;
}

.mb-4 {
    margin-bottom: 1.5rem;
}

.mb-2 {
    margin-bottom: 0.75rem;
}
</style>