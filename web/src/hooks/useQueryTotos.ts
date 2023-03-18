import { todoApi } from "../api-config";
import { useQuery } from "react-query";

export const useQueryTodos = (offset: number) => {
  const getTodos = async () => {
    console.log(offset)
    const data = await todoApi.getTodos({ offset: offset })
    return data
  }

  return useQuery({
    queryKey: "todos",
    queryFn: getTodos,
  })
}
