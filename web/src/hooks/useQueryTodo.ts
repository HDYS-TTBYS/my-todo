import { useQuery } from "react-query";
import appClient from "../app-client";

export const useQueryTodo = (id: number) => {
  const getTodo = async () => {
    return appClient.todo.getTodoId(id)
  }

  return useQuery({
    queryKey: "todo" + id,
    queryFn: getTodo,
  })
}
