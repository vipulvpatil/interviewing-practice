import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Objects;

import edu.princeton.cs.algs4.FlowEdge;
import edu.princeton.cs.algs4.FlowNetwork;
import edu.princeton.cs.algs4.FordFulkerson;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class BaseballElimination {
  private String[] teams;
  private HashMap<String, Integer> teamIndex;

  private int[] wins;
  private int[] losses;
  private int[] remaining;
  private int[][] games;

  private HashMap<String, String[]> eliminators;

  // create a baseball division from given filename in format specified below
  public BaseballElimination(String filename) {
    initStandings(new In(filename));

    for (String team : teams) {
      for (String otherTeam : teams) {
        if (wins(team) + remaining(team) < wins(otherTeam)) {
          String[] eliminatingTeams = { otherTeam };
          eliminators.put(team, eliminatingTeams);
          break;
        }
      }

      if (eliminators.get(team) == null) {
        String[] eliminatingTeams = findEliminatingTeams(team);
        if (eliminatingTeams.length > 0) {
          eliminators.put(team, eliminatingTeams);
        }
      }
    }
  }

  // number of teams
  public int numberOfTeams() {
    return teams.length;
  }

  // all teams
  public Iterable<String> teams() {
    return Arrays.asList(teams);
  }

  // number of wins for given team
  public int wins(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    return wins[teamI];
  }

  // number of losses for given team
  public int losses(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    return losses[teamI];
  }

  // number of remaining games for given team
  public int remaining(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    return remaining[teamI];
  }

  // number of remaining games between team1 and team2
  public int against(String team1, String team2) {
    Integer teamI = teamIndex.get(team1);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    Integer teamJ = teamIndex.get(team2);
    if (teamJ == null) {
      throw new IllegalArgumentException();
    }
    return games[teamI][teamJ];
  }

  // is given team eliminated?
  public boolean isEliminated(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    return eliminators.get(team) != null;
  }

  // subset R of teams that eliminates given team; null if not eliminated
  public Iterable<String> certificateOfElimination(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    String[] cert = eliminators.get(team);
    if (cert == null) {
      return null;
    }
    return Arrays.asList(cert);
  }

  private void initStandings(In in) {
    int noOfTeams = Integer.parseInt(in.readLine());
    teams = new String[noOfTeams];
    teamIndex = new HashMap<>();
    wins = new int[noOfTeams];
    losses = new int[noOfTeams];
    remaining = new int[noOfTeams];
    games = new int[noOfTeams][noOfTeams];
    eliminators = new HashMap<>();
    for (int i = 0; i < noOfTeams; i++) {
      String line = in.readLine();
      line = line.trim();
      String[] tokens = line.split("\\s+");
      String teamName = tokens[0];
      teams[i] = teamName;
      teamIndex.put(teamName, i);
      wins[i] = Integer.parseInt(tokens[1]);
      losses[i] = Integer.parseInt(tokens[2]);
      remaining[i] = Integer.parseInt(tokens[3]);
      for (int j = 0; j < noOfTeams; j++) {
        games[i][j] = Integer.parseInt(tokens[j + 4]);
      }
    }
  }

  private String[] findEliminatingTeams(String team) {
    Integer teamI = teamIndex.get(team);
    if (teamI == null) {
      throw new IllegalArgumentException();
    }
    int teamVerticesCount = teams.length - 1;
    int gameVerticesCount = (teamVerticesCount * (teamVerticesCount - 1)) / 2;
    GameVertexMap gameVertex = new GameVertexMap(teams.length, teamI);
    TeamVertexMap teamVertex = new TeamVertexMap(gameVerticesCount + 1, teamI, teams.length);
    FlowNetwork network = new FlowNetwork(teamVerticesCount + gameVerticesCount + 2);
    for (int i = 0; i < teams.length; i++) {
      if (i != teamI) {
        for (int j = i + 1; j < teams.length; j++) {
          if (j != teamI) {
            int v = gameVertex.indexForGame(i, j);
            network.addEdge(new FlowEdge(0, v, games[i][j]));
            network.addEdge(new FlowEdge(v, teamVertex.indexForTeam(i), Double.MAX_VALUE));
            network.addEdge(new FlowEdge(v, teamVertex.indexForTeam(j), Double.MAX_VALUE));
          }
        }
      }
    }
    for (int i = 0; i < teams.length; i++) {
      if (i != teamI) {
        FlowEdge edge = new FlowEdge(teamVertex.indexForTeam(i), network.V() - 1,
            wins[teamI] + remaining[teamI] - wins[i]);
        network.addEdge(edge);
      }
    }

    FordFulkerson maxflow = new FordFulkerson(network, 0, network.V() - 1);
    ArrayList<String> result = new ArrayList<>();
    for (int i = 0; i < teams.length; i++) {
      if (i != teamI) {
        if (maxflow.inCut(teamVertex.indexForTeam(i))) {
          result.add(teams[i]);
        }
      }
    }

    return result.toArray(new String[0]);
  }

  private class GameVertexMap {
    private final HashMap<Integer, Integer> gameVertex = new HashMap<>();
    private final int teamsLength;

    public GameVertexMap(int teamsLength, int teamI) {
      this.teamsLength = teamsLength;
      int index = 1;
      for (int i = 0; i < teamsLength; i++) {
        if (i != teamI) {
          for (int j = i + 1; j < teamsLength; j++) {
            if (j != teamI) {
              int key = i * teamsLength + j;
              gameVertex.put(key, index++);
            }
          }
        }
      }
    }

    public int indexForGame(int i, int j) {
      int key = i * teamsLength + j;
      Integer result = gameVertex.get(key);
      if (result == null) {
        throw new IllegalArgumentException();
      }
      return result;
    }
  }

  private class TeamVertexMap {
    private final HashMap<Integer, Integer> teamVertex = new HashMap<>();

    public TeamVertexMap(int startIndex, int teamI, int teamsLength) {
      int index = 0;
      for (int i = 0; i < teamsLength; i++) {
        if (i != teamI) {
          int value = startIndex + index;
          teamVertex.put(i, value);
          index++;
        }
      }
    }

    public int indexForTeam(int i) {
      Integer result = teamVertex.get(i);
      if (result == null) {
        throw new IllegalArgumentException();
      }
      return result;
    }
  }

  public static void main(String[] args) {
    BaseballElimination division = new BaseballElimination(args[0]);
    for (String team : division.teams()) {
      if (division.isEliminated(team)) {
        StdOut.print(team + " is eliminated by the subset R = { ");
        for (String t : division.certificateOfElimination(team)) {
          StdOut.print(t + " ");
        }
        StdOut.println("}");
      } else {
        StdOut.println(team + " is not eliminated");
      }
    }

  }
}
