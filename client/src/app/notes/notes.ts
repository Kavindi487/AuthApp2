import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NotesService, Note } from '../services/notes.service';
import { AuthService } from '../services/auth-service';

@Component({
  selector: 'app-notes',
  standalone: true,
  imports: [CommonModule, FormsModule],
  template: `
    <div>
      <h2>Notes</h2>
      <div *ngIf="errorMessage">{{errorMessage}}</div>
      <form (ngSubmit)="createNote()">
        <input [(ngModel)]="title" placeholder="Title" />
        <textarea [(ngModel)]="content" placeholder="Content"></textarea>
        <button type="submit" [disabled]="isLoading">Add Note</button>
      </form>
      <ul>
        <li *ngFor="let note of notes">
          {{note.title}}: {{note.content}}
          <button (click)="deleteNote(note.id)">Delete</button>
        </li>
      </ul>
    </div>
  `
})
export class Notes implements OnInit {
  notes: Note[] = [];
  title = '';
  content = '';
  errorMessage = '';
  isLoading = false;

  constructor(
    private notesService: NotesService,
    public authService: AuthService
  ) {}

  ngOnInit() {
    this.loadNotes();
  }

  loadNotes() {
    this.notesService.getNotes().subscribe({
      next: (res) => this.notes = res.notes,
      error: () => this.errorMessage = 'Could not load notes'
    });
  }

  createNote() {
    if (!this.title || !this.content) return;
    this.isLoading = true;

    this.notesService.createNote(this.title, this.content).subscribe({
      next: () => {
        this.title = '';
        this.content = '';
        this.isLoading = false;
        this.loadNotes();
      },
      error: (err) => {
        this.errorMessage = err.error?.error || 'Failed to create note';
        this.isLoading = false;
      }
    });
  }

  deleteNote(id: number) {
    this.notesService.deleteNote(id).subscribe({
      next: () => this.loadNotes(),
      error: () => this.errorMessage = 'Could not delete note'
    });
  }
}