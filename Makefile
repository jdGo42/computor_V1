# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: jdGo42 <jdGo42@student.42.fr>              +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2020/05/09 00:21:11 by jdGo42            #+#    #+#              #
#    Updated: 2020/10/06 05:37:22 by jdGo42           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #
computorv1

RED=\033[1;31m
GREEN=\033[1;32m
NC=\033[0m

.SILENT:

all: $(NAME)

$(NAME):
	go build src/computorv1.go
	# go build src/*.go
	@printf "$(GREEN)[✓]$(NC) Executable $(NAME) ready!\n"

make:
	go build src/computorv1.go
	# go build src/*.go
	@printf "$(GREEN)[✓]$(NC) Executable $(NAME) ready!\n"

clean:
	@rm -f $(NAME)
	@printf "$(RED)[-]$(NC) Executable $(NAME) deleted\n"

fclean: clean

re: clean all

.PHONY: make re clean all