import React, {useEffect, useState} from 'react'
import PostList from '../components/PostList';
import PostForm from '../components/PostForm';
import MyModal from '../components/UI/modal/MyModal';
import MyButton from '../components/UI/button/MyButton';
import PostService from '../API/PostService';
import Loader from '../components/UI/loader/Loader';
import { useFetching } from '../hooks/useFetching';
import {getPagesCount, getPagesArray} from '../utils/pages';

function Posts() { 
	const [posts, setPosts] = useState([
		{id: 1, title: 'Golang', body: 'Description'},
		{id: 2, title: 'Python', body: 'Description'},
		{id: 3, title: 'Kotlin', body: 'Description'},
	])
	const [modal, setModal] = useState(false);
	const [totalPages, setTotalPages] = useState(0);
	const [limit, setLimit] = useState(10);
	const [page, setPage] = useState(1);

	let pagesArray = getPagesArray(totalPages);

	const [fetchPosts, isPostsLoading, postError] = useFetching(async () => {
		const response = await PostService.getAll(limit, page);
		setPosts(response.data)
		const totalCount = response.headers['x-total-count']
		setTotalPages(getPagesCount(totalCount, limit));
	})

	useEffect(() => {
		fetchPosts()
	}, [page])

	const createPost = (newPost) => {
		setPosts([...posts, newPost])
		setModal(false)
	}

	const removePost = (post) => {
		setPosts(posts.filter(p => p.id !== post.id))
	}

	const changePage = (page) => {
		setPage(page)
	}

	return (
		<div className='App'>
			<MyButton style={{marginTop: 15}} onClick={() => setModal(true)}>Создать пост</MyButton>
			<MyModal visible={modal} setVisible={setModal}>
				<PostForm create={createPost}/>
			</MyModal>

			<hr style={{margin: '15px 0'}}/>

			{postError &&
				<h2>Произошла ошибка $(postError)</h2>
			}

			{isPostsLoading
				? <div style={{display: 'flex', justifyContent: 'center', marginTop: 50}}><Loader/></div>
				: <PostList remove={removePost} posts={posts} />
			}

			<div style={{marginTop: 20, marginBottom: 10}}>
				{pagesArray.map(p =>
					<MyButton onClick={() => changePage(p)}>{p}</MyButton>
				)}
			</div>
		</div>
	);
}

export default Posts;
