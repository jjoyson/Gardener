# Gardener
Backend Docker set up:
1. Set up mongo:
    
    docker run --name some-mongo -d mongo

2. Run tis inside the backend folder: 
    
    docker build .

3. Run the backed service, include flag -d if you dont want to see it run
    
    docker run -p 8000:(port where you want it published) --link some-mongo:mongo (name of image just built)