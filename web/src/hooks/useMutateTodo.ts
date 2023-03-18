import { useQueryClient, useMutation } from "react-query";
import { PostTodoOperationRequest, ToDo, UpdateTodoIdOperationRequest, DeleteTodoIdRequest } from "../types/typescript-fetch";

const fetchE = <T>(url: string, method: "POST" | "PATCH" | "DELETE", json?: any) => {
  let body
  if (!json) {
    body = JSON.stringify(json);
  }
  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json'
  };
  return fetch(url, { method, headers, body }) as Promise<T>
}

export const useMutateTodo = () => {
  const queryClient = useQueryClient()

  const useMutateCreateTodo = useMutation(
    (todo: PostTodoOperationRequest) =>
      fetchE<ToDo>(`${process.env.REACT_APP_REST_URL}/api/todo`, "POST", todo.postTodoRequest!),
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateUpdateTodo = useMutation(
    (todo: UpdateTodoIdOperationRequest) =>
      fetchE<ToDo>(`${process.env.REACT_APP_REST_URL}/api/todo/${todo.id}`, "PATCH", todo.updateTodoIdRequest!),
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  const useMutateDeleteTodo = useMutation(
    (id: DeleteTodoIdRequest) =>
      fetchE<Error>(`${process.env.REACT_APP_REST_URL}/api/todo/${id.id}`, "DELETE"),
    {
      onSuccess: (res: Error) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  return { useMutateCreateTodo, useMutateUpdateTodo, useMutateDeleteTodo }
}
