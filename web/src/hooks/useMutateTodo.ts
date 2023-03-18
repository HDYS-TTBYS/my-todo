import { todoApi } from "../api-config";
import { useQueryClient, useMutation } from "react-query";
import { PostTodoOperationRequest, ToDo } from "../types/typescript-fetch";

export const useMutateTodo = () => {
  const queryClient = useQueryClient()

  const useMutateCreateTodo = useMutation(
    (todo: PostTodoOperationRequest) =>
      todoApi.postTodo(todo),
    {
      onSuccess: (res: ToDo) => {
        queryClient.invalidateQueries('todos')
      },
    }
  )

  return { useMutateCreateTodo }
}
