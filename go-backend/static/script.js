document.addEventListener('DOMContentLoaded', () => {
    loadNotes();
    
    // Handle form submission
    document.getElementById('noteForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const title = document.getElementById('noteTitle').value;
        const content = document.getElementById('noteContent').value;
        
        try {
            const response = await fetch('/api/notes', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ title, content }),
            });
            
            if (response.ok) {
                document.getElementById('noteTitle').value = '';
                document.getElementById('noteContent').value = '';
                loadNotes();
            }
        } catch (error) {
            console.error('Error adding note:', error);
        }
    });
});

async function loadNotes() {
    try {
        const response = await fetch('/api/notes');
        const notes = await response.json();
        displayNotes(notes);
    } catch (error) {
        console.error('Error loading notes:', error);
    }
}

function displayNotes(notes) {
    const notesList = document.getElementById('notesList');
    notesList.innerHTML = '';
    
    notes.forEach(note => {
        const noteElement = document.createElement('div');
        noteElement.className = 'note';
        noteElement.innerHTML = `
            <h3>${note.title}</h3>
            <p>${note.content}</p>
            <div class="note-actions">
                <button onclick="deleteNote('${note.id}')" class="delete-btn">Delete</button>
                <button onclick="editNote('${note.id}')" class="edit-btn">Edit</button>
            </div>
        `;
        notesList.appendChild(noteElement);
    });
}

async function deleteNote(id) {
    if (!confirm('Are you sure you want to delete this note?')) return;
    
    try {
        const response = await fetch(`/api/notes/${id}`, {
            method: 'DELETE',
        });
        
        if (response.ok) {
            loadNotes();
        }
    } catch (error) {
        console.error('Error deleting note:', error);
    }
}

async function editNote(id) {
    const newTitle = prompt('Enter new title:');
    if (!newTitle) return;
    
    const newContent = prompt('Enter new content:');
    if (!newContent) return;
    
    try {
        const response = await fetch(`/api/notes/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: newTitle,
                content: newContent,
            }),
        });
        
        if (response.ok) {
            loadNotes();
        }
    } catch (error) {
        console.error('Error updating note:', error);
    }
}
