import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Note {
  id: number;
  title: string;
  content: string;
  created_at: string;
}

@Injectable({ providedIn: 'root' })
export class NotesService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getNotes(): Observable<{ notes: Note[] }> {
    return this.http.get<{ notes: Note[] }>(`${this.apiUrl}/notes`);
  }

  createNote(title: string, content: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/notes`, { title, content });
  }

  deleteNote(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/notes/${id}`);
  }
}