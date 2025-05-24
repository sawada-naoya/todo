import { fetchTasks } from "../lib/api";
import { useEffect, useState } from "react";
import { Task } from "../types/task";
import { addTask, updateTaskStatus } from "../lib/api";

export default function Home() {
  // tasks 現在のタスクリスト(表示するデータ)
  // settTasks タスクリストを更新する関数
  const [tasks, setTasks] = useState<Task[]>([]);
  // title 入力欄の値
  // setTitle 入力欄の値を更新する関数
  const [title, setTitle] = useState("");

  const loadTasks = async () => {
    const data = await fetchTasks();
    setTasks(data);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // ← フォームの自動リロードを止める
    if (!title.trim()) return; // ← 入力が空なら何もしない
    await addTask(title); // ← 入力されたタイトルをAPI経由で追加
    setTitle(""); // ← 入力欄を空にする
    await loadTasks(); // ← 最新のタスクリストを取得・再描画
  };

  const handleToggle = async (id: number, newDoneStatus: boolean) => {
    try {
      await updateTaskStatus(id, newDoneStatus);
      await loadTasks(); // 再取得して状態反映
    } catch (err) {
      console.error("更新に失敗しました", err);
    }
  };

  // 画面を開いた時にAPIからタスクリストを取得してsetTasksに保存
  useEffect(() => {
    loadTasks();
  }, []);

  return (
    <main className="p-4 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">ToDo List</h1>

      <form onSubmit={handleSubmit} className="flex gap-2 mb-4">
        {/* e.target.value は <input> に入力された値 */}
        {/* それを setTitle(...) に渡して、title 状態に反映 */}
        <input value={title} onChange={(e) => setTitle(e.target.value)} placeholder="新しいタスクを入力" className="border px-2 py-1 flex-1" style={{ padding: "0.5rem 1rem" }} />
        <button type="submit" className="bg-blue-500 text-white px-4 py-1 rounded" style={{ padding: "0.5rem 1rem", marginLeft: 5 }}>
          追加
        </button>
      </form>

      <div className="space-y-2">
        {tasks.map((task) => (
          <div key={task.id} className="flex items-center gap-2 list-none">
            <input type="checkbox" checked={task.is_done} onChange={() => handleToggle(task.id, !task.is_done)} className="cursor-pointer" />
            <span className={task.is_done ? "line-through text-gray-500" : ""}>{task.title}</span>
          </div>
        ))}
      </div>
    </main>
  );
}
