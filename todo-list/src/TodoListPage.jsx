import { useState } from "react";
import { useNavigate } from "react-router-dom";

function TodoListPage() {
    // make it to contain certain character

    const [tasks, setTasks] = useState([
        "Eat breakfast",
        "Take a shower",
        "Walk the dog",
    ]);
    const [newTask, setNewTask] = useState("");
    const navigate = useNavigate();

    function handleInputChange(event) {
        setNewTask(event.target.value);
    }

    function handleLogout(event) {
        navigate("/");
    }

    function addTask() {
        if (newTask.trim() !== "") {
            setTasks((t) => [...t, newTask]);
            setNewTask("");
        }
    }

    function deleteTask(index) {
        const updatedTasks = tasks.filter((_, i) => i !== index);
        setTasks(updatedTasks);
    }

    function moveTaskUp(index) {
        if (index > 0) {
            const updatedTasks = [...tasks];
            [updatedTasks[index], updatedTasks[index - 1]] = [
                updatedTasks[index - 1],
                updatedTasks[index],
            ];
            setTasks(updatedTasks);
        }
    }

    function moveTaskDown(index) {
        if (index < tasks.length - 1) {
            const updatedTasks = [...tasks];
            [updatedTasks[index], updatedTasks[index + 1]] = [
                updatedTasks[index + 1],
                updatedTasks[index],
            ];
            setTasks(updatedTasks);
        }
    }

    return (
        <div className="to-do-list">
            <h1>To-Do-List</h1>
            <div>
                <input
                    type="text"
                    placeholder="Enter a task ..."
                    value={newTask}
                    onChange={handleInputChange}
                />
                <button className="add-button" onClick={addTask}>
                    Add
                </button>
            </div>

            <ol>
                {tasks.map((task, index) => (
                    <li key={index}>
                        <span className="text">{task}</span>
                        <button
                            className="delete-button"
                            onClick={() => deleteTask(index)}
                        >
                            Delete
                        </button>
                        <button
                            className="move-button"
                            onClick={() => moveTaskUp(index)}
                        >
                            👆
                        </button>
                        <button
                            className="move-button"
                            onClick={() => moveTaskDown(index)}
                        >
                            👇
                        </button>
                    </li>
                ))}
            </ol>
            <div className="button-container">
                <button className="normal-button" onClick={handleLogout}>
                    Logout
                </button>
            </div>
        </div>
    );
}

export default TodoListPage;
