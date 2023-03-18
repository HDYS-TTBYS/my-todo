import './App.css';
import { QueryClient, QueryClientProvider } from 'react-query'
import { useQueryTodos } from './hooks/useQueryTotos';
import { ReactQueryDevtools } from 'react-query/devtools'


const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
})

function App() {


  return (
    <QueryClientProvider client={queryClient}>

      <Example />

      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>

  );
}

export default App;

function Example() {
  const { data } = useQueryTodos(0)

  return (
    <div>
      {JSON.stringify(data)}
    </div>
  )
}
