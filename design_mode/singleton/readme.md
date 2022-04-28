# 单例模式

![image](https://user-images.githubusercontent.com/77610866/165486410-0ffcbab5-c250-4010-8916-2f42d1d205d0.png)

单例模式采用了 饿汉式 和 懒汉式 两种实现，个人其实更倾向于饿汉式的实现，简单，并且可以将问题及早暴露，懒汉式虽然支持延迟加载，但是这只是把冷启动时间放到了第一次使用的时候，并没有本质上解决问题，并且为了实现懒汉式还不可避免的需要加锁。


![image](https://user-images.githubusercontent.com/77610866/165727533-5f82b325-c9be-47a7-97a6-a6c5e83f86bd.png)
