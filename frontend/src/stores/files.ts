import { writable } from 'svelte/store';
import type { MusicFile } from '../types';

interface FileState {
  files: MusicFile[];
  loading: boolean;
  error: string | null;
  layout: "grid" | "list";
}

function createFileStore() {
  const initialState: FileState = {
    files: [],
    loading: false,
    error: null,
    layout: "grid"
  };

  const { subscribe, set, update } = writable<FileState>(initialState);

  return {
    subscribe,
    loadFiles: async () => {
      update(state => ({ ...state, loading: true, error: null }));
      
      try {
        // In production, this would be a real API call
        // Simulating API delay
        await new Promise(resolve => setTimeout(resolve, 800));
        
        // Mock data response
        const mockFiles: MusicFile[] = [
          {
            id: '1',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '9',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '10',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '11',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '12',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '13',
            filename: 'bohemian_rhapsody.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 8345678,
            path: '/music/queen/',
            artist: 'Queen',
            album: 'A Night at the Opera',
            title: 'Bohemian Rhapsody',
            year: '1975',
            genre: 'Rock',
            trackNumber: '11',
            bitrate: 320,
            duration: 354,
            sampleRate: 44100
          },
          {
            id: '2',
            filename: 'hotel_california.flac',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'FLAC',
            size: 31457280,
            path: '/music/eagles/',
            artist: 'Eagles',
            album: 'Hotel California',
            title: 'Hotel California',
            year: '1976',
            genre: 'Rock',
            trackNumber: '1',
            bitrate: 1411,
            duration: 391,
            sampleRate: 48000
          },
          {
            id: '3',
            filename: 'billie_jean.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 7234567,
            path: '/music/michael_jackson/',
            artist: 'Michael Jackson',
            album: 'Thriller',
            title: 'Billie Jean',
            year: '1982',
            genre: 'Pop',
            trackNumber: '6',
            bitrate: 320,
            duration: 293,
            sampleRate: 44100
          },
          {
            id: '4',
            filename: 'purple_haze.wav',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'WAV',
            size: 52428800,
            path: '/music/jimi_hendrix/',
            artist: 'Jimi Hendrix',
            album: 'Are You Experienced',
            title: 'Purple Haze',
            year: '1967',
            genre: 'Rock',
            trackNumber: '1',
            bitrate: 1536,
            duration: 167,
            sampleRate: 96000
          },
          {
            id: '5',
            filename: 'sweet_child_o_mine.aac',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'AAC',
            size: 5678912,
            path: '/music/guns_n_roses/',
            artist: 'Guns N\' Roses',
            album: 'Appetite for Destruction',
            title: 'Sweet Child o\' Mine',
            year: '1987',
            genre: 'Hard Rock',
            trackNumber: '9',
            bitrate: 256,
            duration: 356,
            sampleRate: 44100
          },
          {
            id: '6',
            filename: 'smells_like_teen_spirit.mp3',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'MP3',
            size: 7654321,
            path: '/music/nirvana/',
            artist: 'Nirvana',
            album: 'Nevermind',
            title: 'Smells Like Teen Spirit',
            year: '1991',
            genre: 'Grunge',
            trackNumber: '1',
            bitrate: 320,
            duration: 301,
            sampleRate: 44100
          },
          {
            id: '7',
            filename: 'lose_yourself.m4a',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'M4A',
            size: 8765432,
            path: '/music/eminem/',
            artist: 'Eminem',
            album: '8 Mile Soundtrack',
            title: 'Lose Yourself',
            year: '2002',
            genre: 'Hip Hop',
            trackNumber: '1',
            bitrate: 256,
            duration: 326,
            sampleRate: 44100
          },
          {
            id: '8',
            filename: 'comfortably_numb.flac',
            coverUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQkOFXzhaFxI2E7G5iBqHCJ5TxN_-LfnGoYQw&s',
            format: 'FLAC',
            size: 38654321,
            path: '/music/pink_floyd/',
            artist: 'Pink Floyd',
            album: 'The Wall',
            title: 'Comfortably Numb',
            year: '1979',
            genre: 'Progressive Rock',
            trackNumber: '6',
            bitrate: 1411,
            duration: 382,
            sampleRate: 48000
          }
        ];
        
        update(state => ({ ...state, files: mockFiles, loading: false }));
      } catch (error) {
        console.error('Failed to load files:', error);
        update(state => ({ 
          ...state, 
          loading: false,
          error: 'Failed to load music files. Please try again.'
        }));
      }
    },
    updateFile: async (updatedFile: MusicFile) => {
      update(state => {
        // In production, this would be a real API call
        // For now, update the local state
        const updatedFiles = state.files.map(file => 
          file.id === updatedFile.id ? updatedFile : file
        );
        
        return { ...state, files: updatedFiles };
      });
      
      // Simulate API call to update the file
      // In production: await fetch(`/file/${updatedFile.id}`, { method: 'POST', body: JSON.stringify(updatedFile) });
      console.log('File updated:', updatedFile);
    }
  };
}

export const fileStore = createFileStore();
