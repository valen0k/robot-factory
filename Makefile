OBJS = updateStorageCost updateNumberRobots

all:
	go build cmd/updateNumberRobots.go
	go build cmd/updateStorageCost.go

server: all
	go run cmd/robotService.go

#db:
#	psql databasename < init/database.sql

clean:
	rm -f $(OBJS)

fclean:	clean

re:	fclean all

.PHONY:	all clean fclean re