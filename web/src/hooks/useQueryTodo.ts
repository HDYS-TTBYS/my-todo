import { todoApi } from "../api-config";
import { useQuery } from "react-query";

export const useQueryTodo = (id: number) => {
  const getTodo = async () => {
    const data = await todoApi.getTodoId({ id: id })
    return data
  }

  return useQuery({
    queryKey: "todos",
    queryFn: getTodo,
  })
}
