import { create } from 'zustand'

export interface BlogPost {
    id: number,
    name: string,
    contents: string,
}

export interface BlogPostsState {
    blogPosts: BlogPost[]
}

export const useBlogStore = create<BlogPostsState>((set) => ({ blogPosts: [], setBlogPosts: (blogPosts: BlogPost[]) => set({ blogPosts }) }))
