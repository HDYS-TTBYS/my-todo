import { useQueryTodos } from '../hooks/useQueryTotos';

const Index = () => {
  const { data } = useQueryTodos(0)

  return (
    <div>
      {JSON.stringify(data)}
    </div>
  )
}

export default Index
