import { fetchTasks } from "../lib/api";
import { useEffect, useState } from "react";
import { Task } from "../types/task";
import { addTask } from "../lib/api";

export default function Home() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [title, setTitle] = useState("");

  const loadTasks = async () => {
    const data = await fetchTasks();
    setTasks(data);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;
    await addTask(title);
    setTitle("");
    await loadTasks();
  };

  useEffect(() => {
    loadTasks();
  }, []);

  return (
    <main className="p-4 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">ToDo List</h1>

      <form onSubmit={handleSubmit} className="flex gap-2 mb-4">
        <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder="新しいタスクを入力" className="border px-2 py-1 flex-1" />
        <button type="submit" className="bg-blue-500 text-white px-4 py-1 rounded">
          追加
        </button>
      </form>

      <ul className="space-y-2">
        {tasks.map((task) => (
          <li key={task.id} className="border p-2 rounded">
            {task.title}
          </li>
        ))}
      </ul>
    </main>
  );
}
