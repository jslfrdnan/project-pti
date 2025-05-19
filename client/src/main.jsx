import ReactDOM from 'react-dom/client'
import './index.css'

import TodoList from './components/TodoList'
import { TodoProvider } from './contexts/TodoContext'

function App() {
  return (
    <div className="App">
      <TodoProvider>
        <TodoList />
      </TodoProvider>
    </div>
  )
}

ReactDOM.createRoot(document.getElementById('root')).render(<App />)
