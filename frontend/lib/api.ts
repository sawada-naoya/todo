import { Task } from "../types/task";

const API_URL = "http://localhost:8080";

export const fetchTasks = async (): Promise<Task[]> => {
  const res = await fetch(`${API_URL}/tasks`);
  return res.json();
};

export const addTask = async (title: string): Promise<Task> => {
  const res = await fetch(`${API_URL}/tasks`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title }),
  });
  if (!res.ok) {
    throw new Error("Failed to add task");
  }
  const data: Task = await res.json();
  return data;
};

export const updateTaskStatus = async (id: number, is_done: boolean): Promise<void> => {
  await fetch(`${API_URL}/tasks/${id}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ is_done }),
  });
};
